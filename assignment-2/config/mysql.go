package config

import (
	"database/sql"
	"fmt"
)

const (
	host           = "localhost"
	port           = "3306"
	user           = "root"
	password       = "secret"
	dbname         = "assignment_2"
	dbMaxIdle      = 4
	dbMaxOpenConns = 25
)

func ConnectMysql() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxOpenConns)
	db.SetMaxIdleConns(dbMaxIdle)

	return db, nil
}
