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
	dbinfo := "user=" + params.User +
		" password=" + params.Password +
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
