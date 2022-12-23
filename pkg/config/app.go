package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	var err error
	db, err = gorm.Open("postgres", "user=postgres password=0897797817 host=HOST port=5432 dbname=bookstore sslmode=DISABLE")
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
