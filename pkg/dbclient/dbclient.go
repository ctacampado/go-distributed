package dbclient

import (
	"database/sql"
	"os"

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

// NewDB returns an initialized sql.DB pointer
func NewDB() *sql.DB {
	params := &DBParams{
		Addr: DBAddr{
			DBname: os.Getenv("DB_NAME"),
			Host:   os.Getenv("DB_HOST"),
			Port:   os.Getenv("DB_PORT"),
		},
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Sslmode:  os.Getenv("DB_SSL_MODE"),
	}
	return InitDB(params)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
