package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"lobo.tech/task/config"
)

func NewDB(conf *config.Config) *sql.DB {

	str_user := conf.GetValue("local.database_user")
	str_pass := conf.GetValue("local.database_pass")
	str_host := conf.GetValue("local.dbhost")
	str_database := conf.GetValue("local.database")
	connStr := "postgres://" + str_user + ":" + str_pass + "@" + str_host + "/" + str_database + "?sslmode=disable"

	//
	connConfig, _ := pgx.ParseConfig(connStr)
	//	connConfig.Logger = myLogger
	connConfigStr := stdlib.RegisterConnConfig(connConfig)
	//	db, _ := sql.Open("pgx", connStr)
	//
	fmt.Println(connConfigStr)
	db, err := sql.Open("pgx", connConfigStr)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db

}
