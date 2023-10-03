package main

import (
	"log"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	// format "user:password@/dbname"
	db, err := sql.Open("mysql", databaseUrl)

	if err != nil {
		panic(err.Error())
	}

	{
		// create tables if not exists
		createPosts := `
		CREATE TABLE IF NOT EXISTS posts (
			id INT AUTO_INCREMENT,
			title TEXT NOT NULL,
			body TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

		if _, err := db.Exec(createPosts); err != nil {
			log.Fatal(err)
		}

		createUsers := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT,
			name varchar(255) NOT NULL,
			username varchar(255) NOT NULL,
			active boolean,
			PRIMARY KEY (id)
		);`

		if _, err := db.Exec(createUsers); err != nil {
			log.Fatal(err)
		}
	}

	db.Close()
}
