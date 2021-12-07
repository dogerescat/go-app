package models

import (
	"database/sql"
	"log"

	"github.com/dogerescat/go-app/config"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func init() {
	connect()
}

func connect() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
}
