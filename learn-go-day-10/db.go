package main

import (
	"os"

	"database/sql"
)

var DB *sql.DB

// function to initial db connection
func initDB() {
	databaseUrl := os.Getenv("DATABASE_URL")
	// format "user:password@/dbname"
	db, err := sql.Open("mysql", databaseUrl)

	if err != nil {
		panic(err.Error())
	}

	DB = db
}
