package repository

import (
	"database/sql"
	"log"
)

func NewDB() *sql.DB {

	//db, err := sql.Open("mysql", "manager:1234@/cursogo")
	connStr := "postgres://postgres:postgres@localhost/postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db

}
