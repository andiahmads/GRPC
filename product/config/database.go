package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:endi@tcp(localhost:3306)/gogrpc"))
	if err != nil {
		log.Fatalf("Database connection failed %v\n", err.Error())
	}
	return db
}
