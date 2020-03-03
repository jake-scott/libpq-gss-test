package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=wellard.poptart.org dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * from test")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var a string
		if err := rows.Scan(&a); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("a: %s\n", a)
	}
}
