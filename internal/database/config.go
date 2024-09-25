package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetDatabaseConnection() *sql.DB {
	dsn := os.Getenv("DB_URL")
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("unable to connect to database ERR::%v", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatalf("error when get database connection ERR::%v", err)
	}

	return conn
}
