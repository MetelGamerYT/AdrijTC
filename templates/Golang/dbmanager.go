package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func db_connect() {
	cfg := mysql.Config{
		User:                 os.Getenv("UserName"),
		Passwd:               os.Getenv("Password"),
		Net:                  "tcp",
		Addr:                 "IP:Adresse",
		DBName:               "DBName",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database is now connected!")
}

func db_disconnect() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database is now disconnected!")
}

func db_returnquery(query string) ([]string, error) {
	dbquery, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer db_disconnect()

	var result []string
	for dbquery.Next() {
		var rowString string
		if err := dbquery.Scan(&rowString); err != nil {
			return nil, err
		}
		result = append(result, rowString)
	}

	if err = dbquery.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
