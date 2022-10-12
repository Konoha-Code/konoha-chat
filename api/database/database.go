package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "host=127.0.0.1 port=5432 user=USER dbname=chatdb sslmode=disable password=PASSWORD"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB {
	return db
}
