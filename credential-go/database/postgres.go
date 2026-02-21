package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	connStr := "host=ibpr_postgres port=5432 user=ibpr_user password=ibpr_password dbname=middlewareibpr sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	DB = db
	log.Println("PostgreSQL Connected")
}
