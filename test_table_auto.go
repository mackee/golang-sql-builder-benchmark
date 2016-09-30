package db_sql_benchmark

import (
	"strings"
	"strconv"

	"database/sql"
	
	"github.com/mackee/go-sqlla"
)

type testTableSQL struct {
	where sqlla.Where
}

func NewTestTableSQL() testTableSQL {
	q := testTableSQL{}
	return q
}

var testTableAllColumns = []string{
	"b",
}

type testTableSelectSQL struct {
	testTableSQL
	Columns     []string
	order       string
	limit       *uint64
	isForUpdate bool
}

func (q testTableSQL) Select() testTableSelectSQL {
	return testTableSelectSQL{
		q,
		testTableAllColumns,
		"",
		nil,
		false,
	}
}

func (q testTableSelectSQL) Limit(l uint64) testTableSelectSQL {
	q.limit = &l
	return q
}

func (q testTableSelectSQL) ForUpdate() testTableSelectSQL {
	q.isForUpdate = true
	return q
}


func (q testTableSelectSQL) B(v uint32, exprs ...sqlla.Operator) testTableSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "b"}
	q.where = append(q.where, where)
	return q
}

func (q testTableSelectSQL) BIn(vs ...uint32) testTableSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "b"}
	q.where = append(q.where, where)
	return q
}



func (q testTableSelectSQL) OrderByB(order sqlla.Order) testTableSelectSQL {
	q.order = " ORDER BY b"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q testTableSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "SELECT " + columns + " FROM test_table"
	if wheres != "" {
		query += " WHERE" + wheres
	}
	query += q.order
	if q.limit != nil {
		query += " LIMIT " + strconv.FormatUint(*q.limit, 10)
	}

	if q.isForUpdate {
		query += " FOR UPDATE"
	}

	return query + ";", vs, nil
}

func (q testTableSelectSQL) Single(db sqlla.DB) (TestTable, error) {
	q.Columns = testTableAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return TestTable{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q testTableSelectSQL) All(db sqlla.DB) ([]TestTable, error) {
	rs := make([]TestTable, 0, 10)
	q.Columns = testTableAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r, err := q.Scan(rows)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (q testTableSelectSQL) Scan(s sqlla.Scanner) (TestTable, error) {
	var row TestTable
	err := s.Scan(
		&row.B,
		
	)
	return row, err
}

type testTableUpdateSQL struct {
	testTableSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q testTableSQL) Update() testTableUpdateSQL {
	return testTableUpdateSQL{
		testTableSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q testTableUpdateSQL) SetB(v uint32) testTableUpdateSQL {
	q.setMap["b"] = v
	return q
}

func (q testTableUpdateSQL) WhereB(v uint32, exprs ...sqlla.Operator) testTableUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "b"}
	q.where = append(q.where, where)
	return q
}


func (q testTableUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = TestTable{}
	if t, ok := s.(testTableDefaultUpdateHooker); ok {
		q, err = t.DefaultUpdateHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	setColumns, svs, err := q.setMap.ToUpdateSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	wheres, wvs, err := q.where.ToSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "UPDATE test_table SET" + setColumns
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", append(svs, wvs...), nil
}
func (q testTableUpdateSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

type testTableDefaultUpdateHooker interface {
	DefaultUpdateHook(testTableUpdateSQL) (testTableUpdateSQL, error)
}

type testTableInsertSQL struct {
	testTableSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q testTableSQL) Insert() testTableInsertSQL {
	return testTableInsertSQL{
		testTableSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q testTableInsertSQL) ValueB(v uint32) testTableInsertSQL {
	q.setMap["b"] = v
	return q
}


func (q testTableInsertSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = TestTable{}
	if t, ok := s.(testTableDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	qs, vs, err := q.setMap.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "INSERT INTO test_table " + qs

	return query + ";", vs, nil
}

func (q testTableInsertSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		
		return nil, err
	}
	result, err := db.Exec(query, args...)
	return result, err
}

type testTableDefaultInsertHooker interface {
	DefaultInsertHook(testTableInsertSQL) (testTableInsertSQL, error)
}

type testTableDeleteSQL struct {
	testTableSQL
}

func (q testTableSQL) Delete() testTableDeleteSQL {
	return testTableDeleteSQL{
		q,
	}
}


func (q testTableDeleteSQL) B(v uint32, exprs ...sqlla.Operator) testTableDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "b"}
	q.where = append(q.where, where)
	return q
}


func (q testTableDeleteSQL) BIn(vs ...uint32) testTableDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "b"}
	q.where = append(q.where, where)
	return q
}

func (q testTableDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM test_table"
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

func ( q testTableDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

