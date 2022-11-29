package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DbSingleAnswer struct {
	Result string
}

func DBConnect() *sql.DB {

	databaseConfig := ConfigProcess().DatabaseConfig

	db, err := sql.Open("mysql",
		databaseConfig.User+":"+databaseConfig.Password+"@tcp("+databaseConfig.IP+":"+
			databaseConfig.Port+")/"+databaseConfig.DB)

	if err != nil {
		panic(err)
	}

	return db
}

func DBQueryRow(db *sql.DB, query string) DbSingleAnswer {

	result := db.QueryRow(query)

	answer := DbSingleAnswer{}

	err := result.Scan(&answer.Result)

	if err != nil {
		return DbSingleAnswer{Result: "false"}
	}

	return answer
}

func DBQuery(db *sql.DB, query string) bool {
	_, err := db.Query(query)

	if err != nil {
		return false
	} else {
		return true
	}
}
