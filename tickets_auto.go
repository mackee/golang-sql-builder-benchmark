package db_sql_benchmark

import (
	"strings"
	"strconv"

	"database/sql"
	
	"github.com/mackee/go-sqlla"
)

type ticketsSQL struct {
	where sqlla.Where
}

func NewTicketsSQL() ticketsSQL {
	q := ticketsSQL{}
	return q
}

var ticketsAllColumns = []string{
	"id","subdomain_id","state","number",
}

type ticketsSelectSQL struct {
	ticketsSQL
	Columns     []string
	order       string
	limit       *uint64
	isForUpdate bool
}

func (q ticketsSQL) Select() ticketsSelectSQL {
	return ticketsSelectSQL{
		q,
		ticketsAllColumns,
		"",
		nil,
		false,
	}
}

func (q ticketsSelectSQL) Limit(l uint64) ticketsSelectSQL {
	q.limit = &l
	return q
}

func (q ticketsSelectSQL) ForUpdate() ticketsSelectSQL {
	q.isForUpdate = true
	return q
}


func (q ticketsSelectSQL) ID(v uint64, exprs ...sqlla.Operator) ticketsSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint64{Value: v, Op: op, Column: "id"}
	q.where = append(q.where, where)
	return q
}

func (q ticketsSelectSQL) IDIn(vs ...uint64) ticketsSelectSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "id"}
	q.where = append(q.where, where)
	return q
}

func (q ticketsSelectSQL) PkColumn(pk int64, exprs ...sqlla.Operator) ticketsSelectSQL {
	v := uint64(pk)
	return q.ID(v, exprs...)
}

func (q ticketsSelectSQL) OrderByID(order sqlla.Order) ticketsSelectSQL {
	q.order = " ORDER BY id"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q ticketsSelectSQL) SubdomainID(v uint32, exprs ...sqlla.Operator) ticketsSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "subdomain_id"}
	q.where = append(q.where, where)
	return q
}

func (q ticketsSelectSQL) SubdomainIDIn(vs ...uint32) ticketsSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "subdomain_id"}
	q.where = append(q.where, where)
	return q
}



func (q ticketsSelectSQL) OrderBySubdomainID(order sqlla.Order) ticketsSelectSQL {
	q.order = " ORDER BY subdomain_id"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q ticketsSelectSQL) State(v string, exprs ...sqlla.Operator) ticketsSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "state"}
	q.where = append(q.where, where)
	return q
}

func (q ticketsSelectSQL) StateIn(vs ...string) ticketsSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "state"}
	q.where = append(q.where, where)
	return q
}



func (q ticketsSelectSQL) OrderByState(order sqlla.Order) ticketsSelectSQL {
	q.order = " ORDER BY state"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q ticketsSelectSQL) Number(v uint32, exprs ...sqlla.Operator) ticketsSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "number"}
	q.where = append(q.where, where)
	return q
}

func (q ticketsSelectSQL) NumberIn(vs ...uint32) ticketsSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "number"}
	q.where = append(q.where, where)
	return q
}



func (q ticketsSelectSQL) OrderByNumber(order sqlla.Order) ticketsSelectSQL {
	q.order = " ORDER BY number"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q ticketsSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "SELECT " + columns + " FROM tickets"
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

func (s Ticket) Select() (ticketsSelectSQL) {
	return NewTicketsSQL().Select().ID(s.ID)
}
func (q ticketsSelectSQL) Single(db sqlla.DB) (Ticket, error) {
	q.Columns = ticketsAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return Ticket{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q ticketsSelectSQL) All(db sqlla.DB) ([]Ticket, error) {
	rs := make([]Ticket, 0, 10)
	q.Columns = ticketsAllColumns
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

func (q ticketsSelectSQL) Scan(s sqlla.Scanner) (Ticket, error) {
	var row Ticket
	err := s.Scan(
		&row.ID,
		&row.SubdomainID,
		&row.State,
		&row.Number,
		
	)
	return row, err
}

type ticketsUpdateSQL struct {
	ticketsSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q ticketsSQL) Update() ticketsUpdateSQL {
	return ticketsUpdateSQL{
		ticketsSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q ticketsUpdateSQL) SetID(v uint64) ticketsUpdateSQL {
	q.setMap["id"] = v
	return q
}

func (q ticketsUpdateSQL) WhereID(v uint64, exprs ...sqlla.Operator) ticketsUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint64{Value: v, Op: op, Column: "id"}
	q.where = append(q.where, where)
	return q
}


func (q ticketsUpdateSQL) SetSubdomainID(v uint32) ticketsUpdateSQL {
	q.setMap["subdomain_id"] = v
	return q
}

func (q ticketsUpdateSQL) WhereSubdomainID(v uint32, exprs ...sqlla.Operator) ticketsUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "subdomain_id"}
	q.where = append(q.where, where)
	return q
}


func (q ticketsUpdateSQL) SetState(v string) ticketsUpdateSQL {
	q.setMap["state"] = v
	return q
}

func (q ticketsUpdateSQL) WhereState(v string, exprs ...sqlla.Operator) ticketsUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "state"}
	q.where = append(q.where, where)
	return q
}


func (q ticketsUpdateSQL) SetNumber(v uint32) ticketsUpdateSQL {
	q.setMap["number"] = v
	return q
}

func (q ticketsUpdateSQL) WhereNumber(v uint32, exprs ...sqlla.Operator) ticketsUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "number"}
	q.where = append(q.where, where)
	return q
}


func (q ticketsUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = Ticket{}
	if t, ok := s.(ticketsDefaultUpdateHooker); ok {
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

	query := "UPDATE tickets SET" + setColumns
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", append(svs, wvs...), nil
}
func (s Ticket) Update() ticketsUpdateSQL {
	return NewTicketsSQL().Update().WhereID(s.ID)
}

func (q ticketsUpdateSQL) Exec(db sqlla.DB) ([]Ticket, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.ticketsSQL

	return qq.Select().All(db)
}

type ticketsDefaultUpdateHooker interface {
	DefaultUpdateHook(ticketsUpdateSQL) (ticketsUpdateSQL, error)
}

type ticketsInsertSQL struct {
	ticketsSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q ticketsSQL) Insert() ticketsInsertSQL {
	return ticketsInsertSQL{
		ticketsSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q ticketsInsertSQL) ValueID(v uint64) ticketsInsertSQL {
	q.setMap["id"] = v
	return q
}


func (q ticketsInsertSQL) ValueSubdomainID(v uint32) ticketsInsertSQL {
	q.setMap["subdomain_id"] = v
	return q
}


func (q ticketsInsertSQL) ValueState(v string) ticketsInsertSQL {
	q.setMap["state"] = v
	return q
}


func (q ticketsInsertSQL) ValueNumber(v uint32) ticketsInsertSQL {
	q.setMap["number"] = v
	return q
}


func (q ticketsInsertSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = Ticket{}
	if t, ok := s.(ticketsDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	qs, vs, err := q.setMap.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "INSERT INTO tickets " + qs

	return query + ";", vs, nil
}

func (q ticketsInsertSQL) Exec(db sqlla.DB) (Ticket, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return Ticket{}, err
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return Ticket{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Ticket{}, err
	}
	return NewTicketsSQL().Select().PkColumn(id).Single(db)
}

type ticketsDefaultInsertHooker interface {
	DefaultInsertHook(ticketsInsertSQL) (ticketsInsertSQL, error)
}

type ticketsDeleteSQL struct {
	ticketsSQL
}

func (q ticketsSQL) Delete() ticketsDeleteSQL {
	return ticketsDeleteSQL{
		q,
	}
}


func (q ticketsDeleteSQL) ID(v uint64, exprs ...sqlla.Operator) ticketsDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint64{Value: v, Op: op, Column: "id"}
	q.where = append(q.where, where)
	return q
}


func (q ticketsDeleteSQL) IDIn(vs ...uint64) ticketsDeleteSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "id"}
	q.where = append(q.where, where)
	return q
}

func (q ticketsDeleteSQL) SubdomainID(v uint32, exprs ...sqlla.Operator) ticketsDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "subdomain_id"}
	q.where = append(q.where, where)
	return q
}


func (q ticketsDeleteSQL) SubdomainIDIn(vs ...uint32) ticketsDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "subdomain_id"}
	q.where = append(q.where, where)
	return q
}

func (q ticketsDeleteSQL) State(v string, exprs ...sqlla.Operator) ticketsDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "state"}
	q.where = append(q.where, where)
	return q
}


func (q ticketsDeleteSQL) StateIn(vs ...string) ticketsDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "state"}
	q.where = append(q.where, where)
	return q
}

func (q ticketsDeleteSQL) Number(v uint32, exprs ...sqlla.Operator) ticketsDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "number"}
	q.where = append(q.where, where)
	return q
}


func (q ticketsDeleteSQL) NumberIn(vs ...uint32) ticketsDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "number"}
	q.where = append(q.where, where)
	return q
}

func (q ticketsDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM tickets"
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

func ( q ticketsDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}
func (s Ticket) Delete(db sqlla.DB) (sql.Result, error) {
	query, args, err := NewTicketsSQL().Delete().ID(s.ID).ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

