package repository

import (
	"database/sql"
	"log"

	"lobo.tech/task/config"
)

func NewDB(conf *config.Config) *sql.DB {

	str_user := conf.GetValue("local.database_user")
	str_pass := conf.GetValue("local.database_pass")
	str_host := conf.GetValue("local.dbhost")
	str_database := conf.GetValue("local.database")
	connStr := "postgres://" + str_user + ":" + str_pass + "@" + str_host + "/" + str_database
	//fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db.SetMaxIdleConns(5)

	return db

}
