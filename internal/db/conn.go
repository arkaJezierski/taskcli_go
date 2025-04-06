package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"           
	// _ "github.com/mattn/go-sqlite3" 
	_ "github.com/go-sql-driver/mysql" 
)

var DB *sql.DB

func InitDB() {
	dbType := os.Getenv("DB_TYPE") 
	var dsn string

	switch dbType {
		case "postgres":
			dsn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_NAME"),
			)
		// TODO Fix SQLite Error
// 			case "sqlite3":
//		 	dsn = os.Getenv("DB_NAME")
		case "mysql":
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"))
		default:
			log.Fatalf("Unsupported connection type: %s", dbType)
	}

	var err error
	DB, err = sql.Open(dbType, dsn)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

}

