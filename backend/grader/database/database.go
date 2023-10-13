package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Connection *sql.DB = nil

func Init(connectionstring string) error {
	var err error

	Connection, err = sql.Open("mysql", connectionstring)

	if err != nil {
		return err
	}

	return nil
}
