package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:rootpass@tcp(database:3306)/information_schema")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insertVariable, err := db.Query("INSERT INTO learning (`spent_minutes`,`comments`,`work_type`) VALUES (5, 'check some test cases', 'go')")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//insertVariable.Close()

	selectVariable, err := db.Query("SELECT NOW()")

	total := 0

	if err != nil {
		panic(err)
	}

	for selectVariable.Next() {
		selectVariable.Scan(&total)

		fmt.Println(total)
	}

	selectVariable.Close()

}
