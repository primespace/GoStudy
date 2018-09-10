package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	err = db.QueryRow("select aaa from dual").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)
}
