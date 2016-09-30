package db_sql_benchmark

import (
	"strings"
	"strconv"

	"database/sql"
	"time"
	
	"github.com/mackee/go-sqlla"
)

type mytableSQL struct {
	where sqlla.Where
}

func NewMytableSQL() mytableSQL {
	q := mytableSQL{}
	return q
}

var mytableAllColumns = []string{
	"id","a","b","c","price","foo","bar","created","updated",
}

type mytableSelectSQL struct {
	mytableSQL
	Columns     []string
	order       string
	limit       *uint64
	isForUpdate bool
}

func (q mytableSQL) Select() mytableSelectSQL {
	return mytableSelectSQL{
		q,
		mytableAllColumns,
		"",
		nil,
		false,
	}
}

func (q mytableSelectSQL) Limit(l uint64) mytableSelectSQL {
	q.limit = &l
	return q
}

func (q mytableSelectSQL) ForUpdate() mytableSelectSQL {
	q.isForUpdate = true
	return q
}


func (q mytableSelectSQL) ID(v uint32, exprs ...sqlla.Operator) mytableSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "id"}
	q.where = append(q.where, where)
	return q
}

func (q mytableSelectSQL) IDIn(vs ...uint32) mytableSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "id"}
	q.where = append(q.where, where)
	return q
}



func (q mytableSelectSQL) OrderByID(order sqlla.Order) mytableSelectSQL {
	q.order = " ORDER BY id"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q mytableSelectSQL) A(v string, exprs ...sqlla.Operator) mytableSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "a"}
	q.where = append(q.where, where)
	return q
}

func (q mytableSelectSQL) AIn(vs ...string) mytableSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "a"}
	q.where = append(q.where, where)
	return q
}



func (q mytableSelectSQL) OrderByA(order sqlla.Order) mytableSelectSQL {
	q.order = " ORDER BY a"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q mytableSelectSQL) B(v string, exprs ...sqlla.Operator) mytableSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "b"}
	q.where = append(q.where, where)
	return q
}

func (q mytableSelectSQL) BIn(vs ...string) mytableSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "b"}
	q.where = append(q.where, where)
	return q
}



func (q mytableSelectSQL) OrderByB(order sqlla.Order) mytableSelectSQL {
	q.order = " ORDER BY b"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q mytableSelectSQL) C(v uint32, exprs ...sqlla.Operator) mytableSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "c"}
	q.where = append(q.where, where)
	return q
}

func (q mytableSelectSQL) CIn(vs ...uint32) mytableSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "c"}
	q.where = append(q.where, where)
	return q
}



func (q mytableSelectSQL) OrderByC(order sqlla.Order) mytableSelectSQL {
	q.order = " ORDER BY c"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q mytableSelectSQL) Price(v float64, exprs ...sqlla.Operator) mytableSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprFloat64{Value: v, Op: op, Column: "price"}
	q.where = append(q.where, where)
	return q
}

func (q mytableSelectSQL) PriceIn(vs ...float64) mytableSelectSQL {
	where := sqlla.ExprMultiFloat64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "price"}
	q.where = append(q.where, where)
	return q
}



func (q mytableSelectSQL) OrderByPrice(order sqlla.Order) mytableSelectSQL {
	q.order = " ORDER BY price"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q mytableSelectSQL) Foo(v uint32, exprs ...sqlla.Operator) mytableSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "foo"}
	q.where = append(q.where, where)
	return q
}

func (q mytableSelectSQL) FooIn(vs ...uint32) mytableSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "foo"}
	q.where = append(q.where, where)
	return q
}



func (q mytableSelectSQL) OrderByFoo(order sqlla.Order) mytableSelectSQL {
	q.order = " ORDER BY foo"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q mytableSelectSQL) Bar(v sql.NullInt64, exprs ...sqlla.Operator) mytableSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullInt64{Value: v, Op: op, Column: "bar"}
	q.where = append(q.where, where)
	return q
}

func (q mytableSelectSQL) BarIn(vs ...sql.NullInt64) mytableSelectSQL {
	where := sqlla.ExprMultiNullInt64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "bar"}
	q.where = append(q.where, where)
	return q
}



func (q mytableSelectSQL) OrderByBar(order sqlla.Order) mytableSelectSQL {
	q.order = " ORDER BY bar"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q mytableSelectSQL) Created(v time.Time, exprs ...sqlla.Operator) mytableSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "created"}
	q.where = append(q.where, where)
	return q
}

func (q mytableSelectSQL) CreatedIn(vs ...time.Time) mytableSelectSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "created"}
	q.where = append(q.where, where)
	return q
}



func (q mytableSelectSQL) OrderByCreated(order sqlla.Order) mytableSelectSQL {
	q.order = " ORDER BY created"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q mytableSelectSQL) Updated(v time.Time, exprs ...sqlla.Operator) mytableSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "updated"}
	q.where = append(q.where, where)
	return q
}

func (q mytableSelectSQL) UpdatedIn(vs ...time.Time) mytableSelectSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "updated"}
	q.where = append(q.where, where)
	return q
}



func (q mytableSelectSQL) OrderByUpdated(order sqlla.Order) mytableSelectSQL {
	q.order = " ORDER BY updated"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q mytableSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "SELECT " + columns + " FROM mytable"
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

func (q mytableSelectSQL) Single(db sqlla.DB) (Mytable, error) {
	q.Columns = mytableAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return Mytable{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q mytableSelectSQL) All(db sqlla.DB) ([]Mytable, error) {
	rs := make([]Mytable, 0, 10)
	q.Columns = mytableAllColumns
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

func (q mytableSelectSQL) Scan(s sqlla.Scanner) (Mytable, error) {
	var row Mytable
	err := s.Scan(
		&row.ID,
		&row.A,
		&row.B,
		&row.C,
		&row.Price,
		&row.Foo,
		&row.Bar,
		&row.Created,
		&row.Updated,
		
	)
	return row, err
}

type mytableUpdateSQL struct {
	mytableSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q mytableSQL) Update() mytableUpdateSQL {
	return mytableUpdateSQL{
		mytableSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q mytableUpdateSQL) SetID(v uint32) mytableUpdateSQL {
	q.setMap["id"] = v
	return q
}

func (q mytableUpdateSQL) WhereID(v uint32, exprs ...sqlla.Operator) mytableUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "id"}
	q.where = append(q.where, where)
	return q
}


func (q mytableUpdateSQL) SetA(v string) mytableUpdateSQL {
	q.setMap["a"] = v
	return q
}

func (q mytableUpdateSQL) WhereA(v string, exprs ...sqlla.Operator) mytableUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "a"}
	q.where = append(q.where, where)
	return q
}


func (q mytableUpdateSQL) SetB(v string) mytableUpdateSQL {
	q.setMap["b"] = v
	return q
}

func (q mytableUpdateSQL) WhereB(v string, exprs ...sqlla.Operator) mytableUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "b"}
	q.where = append(q.where, where)
	return q
}


func (q mytableUpdateSQL) SetC(v uint32) mytableUpdateSQL {
	q.setMap["c"] = v
	return q
}

func (q mytableUpdateSQL) WhereC(v uint32, exprs ...sqlla.Operator) mytableUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "c"}
	q.where = append(q.where, where)
	return q
}


func (q mytableUpdateSQL) SetPrice(v float64) mytableUpdateSQL {
	q.setMap["price"] = v
	return q
}

func (q mytableUpdateSQL) WherePrice(v float64, exprs ...sqlla.Operator) mytableUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprFloat64{Value: v, Op: op, Column: "price"}
	q.where = append(q.where, where)
	return q
}


func (q mytableUpdateSQL) SetFoo(v uint32) mytableUpdateSQL {
	q.setMap["foo"] = v
	return q
}

func (q mytableUpdateSQL) WhereFoo(v uint32, exprs ...sqlla.Operator) mytableUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "foo"}
	q.where = append(q.where, where)
	return q
}


func (q mytableUpdateSQL) SetBar(v sql.NullInt64) mytableUpdateSQL {
	q.setMap["bar"] = v
	return q
}

func (q mytableUpdateSQL) WhereBar(v sql.NullInt64, exprs ...sqlla.Operator) mytableUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullInt64{Value: v, Op: op, Column: "bar"}
	q.where = append(q.where, where)
	return q
}


func (q mytableUpdateSQL) SetCreated(v time.Time) mytableUpdateSQL {
	q.setMap["created"] = v
	return q
}

func (q mytableUpdateSQL) WhereCreated(v time.Time, exprs ...sqlla.Operator) mytableUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "created"}
	q.where = append(q.where, where)
	return q
}


func (q mytableUpdateSQL) SetUpdated(v time.Time) mytableUpdateSQL {
	q.setMap["updated"] = v
	return q
}

func (q mytableUpdateSQL) WhereUpdated(v time.Time, exprs ...sqlla.Operator) mytableUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "updated"}
	q.where = append(q.where, where)
	return q
}


func (q mytableUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = Mytable{}
	if t, ok := s.(mytableDefaultUpdateHooker); ok {
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

	query := "UPDATE mytable SET" + setColumns
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", append(svs, wvs...), nil
}
func (q mytableUpdateSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

type mytableDefaultUpdateHooker interface {
	DefaultUpdateHook(mytableUpdateSQL) (mytableUpdateSQL, error)
}

type mytableInsertSQL struct {
	mytableSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q mytableSQL) Insert() mytableInsertSQL {
	return mytableInsertSQL{
		mytableSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q mytableInsertSQL) ValueID(v uint32) mytableInsertSQL {
	q.setMap["id"] = v
	return q
}


func (q mytableInsertSQL) ValueA(v string) mytableInsertSQL {
	q.setMap["a"] = v
	return q
}


func (q mytableInsertSQL) ValueB(v string) mytableInsertSQL {
	q.setMap["b"] = v
	return q
}


func (q mytableInsertSQL) ValueC(v uint32) mytableInsertSQL {
	q.setMap["c"] = v
	return q
}


func (q mytableInsertSQL) ValuePrice(v float64) mytableInsertSQL {
	q.setMap["price"] = v
	return q
}


func (q mytableInsertSQL) ValueFoo(v uint32) mytableInsertSQL {
	q.setMap["foo"] = v
	return q
}


func (q mytableInsertSQL) ValueBar(v sql.NullInt64) mytableInsertSQL {
	q.setMap["bar"] = v
	return q
}


func (q mytableInsertSQL) ValueCreated(v time.Time) mytableInsertSQL {
	q.setMap["created"] = v
	return q
}


func (q mytableInsertSQL) ValueUpdated(v time.Time) mytableInsertSQL {
	q.setMap["updated"] = v
	return q
}


func (q mytableInsertSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = Mytable{}
	if t, ok := s.(mytableDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	qs, vs, err := q.setMap.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "INSERT INTO mytable " + qs

	return query + ";", vs, nil
}

func (q mytableInsertSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		
		return nil, err
	}
	result, err := db.Exec(query, args...)
	return result, err
}

type mytableDefaultInsertHooker interface {
	DefaultInsertHook(mytableInsertSQL) (mytableInsertSQL, error)
}

type mytableDeleteSQL struct {
	mytableSQL
}

func (q mytableSQL) Delete() mytableDeleteSQL {
	return mytableDeleteSQL{
		q,
	}
}


func (q mytableDeleteSQL) ID(v uint32, exprs ...sqlla.Operator) mytableDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "id"}
	q.where = append(q.where, where)
	return q
}


func (q mytableDeleteSQL) IDIn(vs ...uint32) mytableDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "id"}
	q.where = append(q.where, where)
	return q
}

func (q mytableDeleteSQL) A(v string, exprs ...sqlla.Operator) mytableDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "a"}
	q.where = append(q.where, where)
	return q
}


func (q mytableDeleteSQL) AIn(vs ...string) mytableDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "a"}
	q.where = append(q.where, where)
	return q
}

func (q mytableDeleteSQL) B(v string, exprs ...sqlla.Operator) mytableDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "b"}
	q.where = append(q.where, where)
	return q
}


func (q mytableDeleteSQL) BIn(vs ...string) mytableDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "b"}
	q.where = append(q.where, where)
	return q
}

func (q mytableDeleteSQL) C(v uint32, exprs ...sqlla.Operator) mytableDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "c"}
	q.where = append(q.where, where)
	return q
}


func (q mytableDeleteSQL) CIn(vs ...uint32) mytableDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "c"}
	q.where = append(q.where, where)
	return q
}

func (q mytableDeleteSQL) Price(v float64, exprs ...sqlla.Operator) mytableDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprFloat64{Value: v, Op: op, Column: "price"}
	q.where = append(q.where, where)
	return q
}


func (q mytableDeleteSQL) PriceIn(vs ...float64) mytableDeleteSQL {
	where := sqlla.ExprMultiFloat64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "price"}
	q.where = append(q.where, where)
	return q
}

func (q mytableDeleteSQL) Foo(v uint32, exprs ...sqlla.Operator) mytableDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "foo"}
	q.where = append(q.where, where)
	return q
}


func (q mytableDeleteSQL) FooIn(vs ...uint32) mytableDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "foo"}
	q.where = append(q.where, where)
	return q
}

func (q mytableDeleteSQL) Bar(v sql.NullInt64, exprs ...sqlla.Operator) mytableDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullInt64{Value: v, Op: op, Column: "bar"}
	q.where = append(q.where, where)
	return q
}


func (q mytableDeleteSQL) BarIn(vs ...sql.NullInt64) mytableDeleteSQL {
	where := sqlla.ExprMultiNullInt64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "bar"}
	q.where = append(q.where, where)
	return q
}

func (q mytableDeleteSQL) Created(v time.Time, exprs ...sqlla.Operator) mytableDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "created"}
	q.where = append(q.where, where)
	return q
}


func (q mytableDeleteSQL) CreatedIn(vs ...time.Time) mytableDeleteSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "created"}
	q.where = append(q.where, where)
	return q
}

func (q mytableDeleteSQL) Updated(v time.Time, exprs ...sqlla.Operator) mytableDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "updated"}
	q.where = append(q.where, where)
	return q
}


func (q mytableDeleteSQL) UpdatedIn(vs ...time.Time) mytableDeleteSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "updated"}
	q.where = append(q.where, where)
	return q
}

func (q mytableDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM mytable"
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

func ( q mytableDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

