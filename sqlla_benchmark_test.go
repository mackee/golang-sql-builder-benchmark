package db_sql_benchmark

import (
	"database/sql"
	"testing"
	"time"
)

//go:generate sqlla

//+table: tickets
type Ticket struct {
	ID          uint64 `db:"id,primarykey"`
	SubdomainID uint32 `db:"subdomain_id"`
	State       string `db:"state"`
	Number      uint32 `db:"number"`
}

//+table: c
type C struct {
	A   string `db:"a"`
	B   string `db:"b"`
	C   string `db:"c"`
	D   uint32 `db:"d"`
	E   string `db:"e"`
	F   uint32 `db:"f"`
	G   uint32 `db:"g"`
	H   uint32 `db:"h"`
	I   uint32 `db:"i"`
	II  uint32 `db:"ii"`
	III uint32 `db:"iii"`
	J   uint32 `db:"j"`
	JJ  uint32 `db:"jj"`
	JJJ uint32 `db:"jjj"`
	K   uint32 `db:"k"`
	l   uint32 `db:"l"`
	Z   string `db:"z"`
	Y   string `db:"y"`
	X   string `db:"x"`
}

//+table: mytable
type Mytable struct {
	ID      uint32        `db:"id"`
	A       string        `db:"a"`
	B       string        `db:"b"`
	C       uint32        `db:"c"`
	Price   float64       `db:"price"`
	Foo     uint32        `db:"foo"`
	Bar     sql.NullInt64 `db:"bar"`
	Created time.Time     `db:"created"`
	Updated time.Time     `db:"updated"`
}

//+table: test_table
type TestTable struct {
	B uint32 `db:"b"`
}

func BenchmarkSqllaSelectSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sql := NewTicketsSQL().Select()
		sql.Columns = []string{"id"}
		sql.SubdomainID(uint32(1)).StateIn("open", "span").ToSql()
	}
}

func BenchmarkSqllaInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewMytableSQL().Insert().
			ValueID(uint32(1)).
			ValueA("test_a").
			ValueB("test_b").
			ValuePrice(100.05).
			ValueCreated(time.Date(2014, time.January, 5, 0, 0, 0, 0, time.UTC)).
			ValueUpdated(time.Date(2015, time.January, 5, 0, 0, 0, 0, time.UTC)).
			ToSql()
	}
}
