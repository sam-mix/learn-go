package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Fl-GtuH3*OJOrOFQCvdZb5LXX2giZB@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err)
	}
	rs, err := db.Exec("show databases")
	if err != nil {
		panic(err)
	}
	fmt.Printf("rs: %v\n", rs)
}
