package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func initDatabase() {
	var err error
	if testingMode { // you can set this variable based on a flag or environment variable
		Db, err = sql.Open("sqlite3", ":memory:")
	} else {
		Db, err = sql.Open("sqlite3", "./todoapi.db")
	}
	if err != nil {
		panic(err)
	}

	Db, err = sql.Open("sqlite3", "./todoapi.db")
	if err != nil {
		panic(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS todos (
    	id INTEGER PRIMARY KEY,
    	title TEXT,
    	done BOOLEAN
	);`

	_, err = Db.Exec(createTable)
	if err != nil {
		panic(err)
	}
}
