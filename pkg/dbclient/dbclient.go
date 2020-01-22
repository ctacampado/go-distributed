package dbclient

import (
	"database/sql"

	_ "github.com/lib/pq" //pq
)

//DBAddr -
type DBAddr struct {
	DBname string
	Host   string
	Port   string
}

//DBParams -
type DBParams struct {
	User     string
	Password string
	Addr     DBAddr
	Sslmode  string
}

//InitDB -
func InitDB(params *DBParams) *sql.DB {
	var dbinfo string
	if 0 != len(params.Password) {
		dbinfo += "user=" + params.Password
	}

	dbinfo += "user=" + params.User +
		" dbname=" + params.Addr.DBname +
		" sslmode=" + params.Sslmode +
		" port=" + params.Addr.Port

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

// NewDB returns an initialized sql.DB pointer
func NewDB(p *DBParams) *sql.DB {
	params := p
	return InitDB(params)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
