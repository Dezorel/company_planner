package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "dezorel:dezorelpass@tcp(database:3306)/information_schema")

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

	result := ""

	if err != nil {
		panic(err)
	}

	for selectVariable.Next() {
		selectVariable.Scan(&result)

		fmt.Println(result)
	}

	selectVariable.Close()

}
