package db_sql_benchmark

import (
	"strings"
	"strconv"

	"database/sql"
	
	"github.com/mackee/go-sqlla"
)

type cSQL struct {
	where sqlla.Where
}

func NewCSQL() cSQL {
	q := cSQL{}
	return q
}

var cAllColumns = []string{
	"a","b","c","d","e","f","g","h","i","ii","iii","j","jj","jjj","k","l","z","y","x",
}

type cSelectSQL struct {
	cSQL
	Columns     []string
	order       string
	limit       *uint64
	isForUpdate bool
}

func (q cSQL) Select() cSelectSQL {
	return cSelectSQL{
		q,
		cAllColumns,
		"",
		nil,
		false,
	}
}

func (q cSelectSQL) Limit(l uint64) cSelectSQL {
	q.limit = &l
	return q
}

func (q cSelectSQL) ForUpdate() cSelectSQL {
	q.isForUpdate = true
	return q
}


func (q cSelectSQL) A(v string, exprs ...sqlla.Operator) cSelectSQL {
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

func (q cSelectSQL) AIn(vs ...string) cSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "a"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByA(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY a"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) B(v string, exprs ...sqlla.Operator) cSelectSQL {
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

func (q cSelectSQL) BIn(vs ...string) cSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "b"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByB(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY b"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) C(v string, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "c"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) CIn(vs ...string) cSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "c"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByC(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY c"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) D(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "d"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) DIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "d"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByD(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY d"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) E(v string, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "e"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) EIn(vs ...string) cSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "e"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByE(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY e"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) F(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "f"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) FIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "f"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByF(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY f"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) G(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "g"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) GIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "g"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByG(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY g"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) H(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "h"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) HIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "h"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByH(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY h"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) I(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "i"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) IIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "i"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByI(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY i"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) Ii(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "ii"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) IiIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "ii"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByIi(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY ii"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) Iii(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "iii"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) IiiIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "iii"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByIii(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY iii"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) J(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "j"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) JIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "j"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByJ(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY j"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) Jj(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "jj"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) JjIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "jj"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByJj(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY jj"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) Jjj(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "jjj"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) JjjIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "jjj"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByJjj(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY jjj"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) K(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "k"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) KIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "k"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByK(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY k"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) L(v uint32, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "l"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) LIn(vs ...uint32) cSelectSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "l"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByL(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY l"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) Z(v string, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "z"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) ZIn(vs ...string) cSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "z"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByZ(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY z"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) Y(v string, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "y"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) YIn(vs ...string) cSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "y"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByY(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY y"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) X(v string, exprs ...sqlla.Operator) cSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "x"}
	q.where = append(q.where, where)
	return q
}

func (q cSelectSQL) XIn(vs ...string) cSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "x"}
	q.where = append(q.where, where)
	return q
}



func (q cSelectSQL) OrderByX(order sqlla.Order) cSelectSQL {
	q.order = " ORDER BY x"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q cSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "SELECT " + columns + " FROM c"
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

func (q cSelectSQL) Single(db sqlla.DB) (C, error) {
	q.Columns = cAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return C{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q cSelectSQL) All(db sqlla.DB) ([]C, error) {
	rs := make([]C, 0, 10)
	q.Columns = cAllColumns
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

func (q cSelectSQL) Scan(s sqlla.Scanner) (C, error) {
	var row C
	err := s.Scan(
		&row.A,
		&row.B,
		&row.C,
		&row.D,
		&row.E,
		&row.F,
		&row.G,
		&row.H,
		&row.I,
		&row.II,
		&row.III,
		&row.J,
		&row.JJ,
		&row.JJJ,
		&row.K,
		&row.l,
		&row.Z,
		&row.Y,
		&row.X,
		
	)
	return row, err
}

type cUpdateSQL struct {
	cSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q cSQL) Update() cUpdateSQL {
	return cUpdateSQL{
		cSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q cUpdateSQL) SetA(v string) cUpdateSQL {
	q.setMap["a"] = v
	return q
}

func (q cUpdateSQL) WhereA(v string, exprs ...sqlla.Operator) cUpdateSQL {
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


func (q cUpdateSQL) SetB(v string) cUpdateSQL {
	q.setMap["b"] = v
	return q
}

func (q cUpdateSQL) WhereB(v string, exprs ...sqlla.Operator) cUpdateSQL {
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


func (q cUpdateSQL) SetC(v string) cUpdateSQL {
	q.setMap["c"] = v
	return q
}

func (q cUpdateSQL) WhereC(v string, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "c"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetD(v uint32) cUpdateSQL {
	q.setMap["d"] = v
	return q
}

func (q cUpdateSQL) WhereD(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "d"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetE(v string) cUpdateSQL {
	q.setMap["e"] = v
	return q
}

func (q cUpdateSQL) WhereE(v string, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "e"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetF(v uint32) cUpdateSQL {
	q.setMap["f"] = v
	return q
}

func (q cUpdateSQL) WhereF(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "f"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetG(v uint32) cUpdateSQL {
	q.setMap["g"] = v
	return q
}

func (q cUpdateSQL) WhereG(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "g"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetH(v uint32) cUpdateSQL {
	q.setMap["h"] = v
	return q
}

func (q cUpdateSQL) WhereH(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "h"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetI(v uint32) cUpdateSQL {
	q.setMap["i"] = v
	return q
}

func (q cUpdateSQL) WhereI(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "i"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetIi(v uint32) cUpdateSQL {
	q.setMap["ii"] = v
	return q
}

func (q cUpdateSQL) WhereIi(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "ii"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetIii(v uint32) cUpdateSQL {
	q.setMap["iii"] = v
	return q
}

func (q cUpdateSQL) WhereIii(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "iii"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetJ(v uint32) cUpdateSQL {
	q.setMap["j"] = v
	return q
}

func (q cUpdateSQL) WhereJ(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "j"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetJj(v uint32) cUpdateSQL {
	q.setMap["jj"] = v
	return q
}

func (q cUpdateSQL) WhereJj(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "jj"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetJjj(v uint32) cUpdateSQL {
	q.setMap["jjj"] = v
	return q
}

func (q cUpdateSQL) WhereJjj(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "jjj"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetK(v uint32) cUpdateSQL {
	q.setMap["k"] = v
	return q
}

func (q cUpdateSQL) WhereK(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "k"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetL(v uint32) cUpdateSQL {
	q.setMap["l"] = v
	return q
}

func (q cUpdateSQL) WhereL(v uint32, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "l"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetZ(v string) cUpdateSQL {
	q.setMap["z"] = v
	return q
}

func (q cUpdateSQL) WhereZ(v string, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "z"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetY(v string) cUpdateSQL {
	q.setMap["y"] = v
	return q
}

func (q cUpdateSQL) WhereY(v string, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "y"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) SetX(v string) cUpdateSQL {
	q.setMap["x"] = v
	return q
}

func (q cUpdateSQL) WhereX(v string, exprs ...sqlla.Operator) cUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "x"}
	q.where = append(q.where, where)
	return q
}


func (q cUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = C{}
	if t, ok := s.(cDefaultUpdateHooker); ok {
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

	query := "UPDATE c SET" + setColumns
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", append(svs, wvs...), nil
}
func (q cUpdateSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

type cDefaultUpdateHooker interface {
	DefaultUpdateHook(cUpdateSQL) (cUpdateSQL, error)
}

type cInsertSQL struct {
	cSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q cSQL) Insert() cInsertSQL {
	return cInsertSQL{
		cSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q cInsertSQL) ValueA(v string) cInsertSQL {
	q.setMap["a"] = v
	return q
}


func (q cInsertSQL) ValueB(v string) cInsertSQL {
	q.setMap["b"] = v
	return q
}


func (q cInsertSQL) ValueC(v string) cInsertSQL {
	q.setMap["c"] = v
	return q
}


func (q cInsertSQL) ValueD(v uint32) cInsertSQL {
	q.setMap["d"] = v
	return q
}


func (q cInsertSQL) ValueE(v string) cInsertSQL {
	q.setMap["e"] = v
	return q
}


func (q cInsertSQL) ValueF(v uint32) cInsertSQL {
	q.setMap["f"] = v
	return q
}


func (q cInsertSQL) ValueG(v uint32) cInsertSQL {
	q.setMap["g"] = v
	return q
}


func (q cInsertSQL) ValueH(v uint32) cInsertSQL {
	q.setMap["h"] = v
	return q
}


func (q cInsertSQL) ValueI(v uint32) cInsertSQL {
	q.setMap["i"] = v
	return q
}


func (q cInsertSQL) ValueIi(v uint32) cInsertSQL {
	q.setMap["ii"] = v
	return q
}


func (q cInsertSQL) ValueIii(v uint32) cInsertSQL {
	q.setMap["iii"] = v
	return q
}


func (q cInsertSQL) ValueJ(v uint32) cInsertSQL {
	q.setMap["j"] = v
	return q
}


func (q cInsertSQL) ValueJj(v uint32) cInsertSQL {
	q.setMap["jj"] = v
	return q
}


func (q cInsertSQL) ValueJjj(v uint32) cInsertSQL {
	q.setMap["jjj"] = v
	return q
}


func (q cInsertSQL) ValueK(v uint32) cInsertSQL {
	q.setMap["k"] = v
	return q
}


func (q cInsertSQL) ValueL(v uint32) cInsertSQL {
	q.setMap["l"] = v
	return q
}


func (q cInsertSQL) ValueZ(v string) cInsertSQL {
	q.setMap["z"] = v
	return q
}


func (q cInsertSQL) ValueY(v string) cInsertSQL {
	q.setMap["y"] = v
	return q
}


func (q cInsertSQL) ValueX(v string) cInsertSQL {
	q.setMap["x"] = v
	return q
}


func (q cInsertSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = C{}
	if t, ok := s.(cDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	qs, vs, err := q.setMap.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "INSERT INTO c " + qs

	return query + ";", vs, nil
}

func (q cInsertSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		
		return nil, err
	}
	result, err := db.Exec(query, args...)
	return result, err
}

type cDefaultInsertHooker interface {
	DefaultInsertHook(cInsertSQL) (cInsertSQL, error)
}

type cDeleteSQL struct {
	cSQL
}

func (q cSQL) Delete() cDeleteSQL {
	return cDeleteSQL{
		q,
	}
}


func (q cDeleteSQL) A(v string, exprs ...sqlla.Operator) cDeleteSQL {
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


func (q cDeleteSQL) AIn(vs ...string) cDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "a"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) B(v string, exprs ...sqlla.Operator) cDeleteSQL {
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


func (q cDeleteSQL) BIn(vs ...string) cDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "b"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) C(v string, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "c"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) CIn(vs ...string) cDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "c"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) D(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "d"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) DIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "d"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) E(v string, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "e"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) EIn(vs ...string) cDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "e"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) F(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "f"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) FIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "f"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) G(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "g"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) GIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "g"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) H(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "h"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) HIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "h"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) I(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "i"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) IIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "i"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) Ii(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "ii"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) IiIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "ii"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) Iii(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "iii"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) IiiIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "iii"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) J(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "j"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) JIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "j"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) Jj(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "jj"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) JjIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "jj"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) Jjj(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "jjj"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) JjjIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "jjj"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) K(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "k"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) KIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "k"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) L(v uint32, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint32{Value: v, Op: op, Column: "l"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) LIn(vs ...uint32) cDeleteSQL {
	where := sqlla.ExprMultiUint32{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "l"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) Z(v string, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "z"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) ZIn(vs ...string) cDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "z"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) Y(v string, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "y"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) YIn(vs ...string) cDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "y"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) X(v string, exprs ...sqlla.Operator) cDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "x"}
	q.where = append(q.where, where)
	return q
}


func (q cDeleteSQL) XIn(vs ...string) cDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "x"}
	q.where = append(q.where, where)
	return q
}

func (q cDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM c"
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

func ( q cDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

