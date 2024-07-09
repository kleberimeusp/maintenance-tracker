package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/maintenance_tasks")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
