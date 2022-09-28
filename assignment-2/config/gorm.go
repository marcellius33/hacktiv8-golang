package config

import (
	"assignment-2/httpserver/repositories/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

func ConnectMysqlGORM() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})

	return db, nil
}
