package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// MustConnectToDB - creates connection to database, or panic if connection not established
func MustConnectToDB() *sql.DB {
	DSN := "host=localhost port=15432 user=postgres password=postgres dbname=pet_store sslmode=disable"
	db, err := sql.Open("postgres", DSN)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
