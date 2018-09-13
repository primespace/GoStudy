package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:@tcp(localhost)/chartschool")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select nickname from user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("columns ", len(columns))

	for rows.Next() {
		rows.Scan
	}

	// var name string
	// err = db.QueryRow("select nickname from user").Scan(&name)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(name)
}
