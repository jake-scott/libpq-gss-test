package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("host=%s dbname=test sslmode=disable", hn)
	fmt.Printf("Connect string: %s\n", connStr)
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
