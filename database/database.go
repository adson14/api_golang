package database

import (
	"api_golang/database/migrations"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb() {
	str := "host=localhost user=postgres password=123456 dbname=api_go port=5432 sslmode=disable"

	database, erro := gorm.Open(postgres.Open(str), &gorm.Config{})

	if erro != nil {
		log.Fatal("error: ", erro)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RodaMigration(db)

}

func GetDatabase() *gorm.DB {
	return db
}
