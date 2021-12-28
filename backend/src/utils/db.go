package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/miello/url-shortener/backend/src/dto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	var err error
	DB_USER := os.Getenv("DATABASE_USER")
	if DB_USER == "" {
		DB_USER = "postgres"
	}

	DB_PASS := os.Getenv("DATABASE_PASSWORD")
	if DB_PASS == "" {
		DB_PASS = "example"
	}

	DB_NAME := os.Getenv("DATABASE_NAME")
	if DB_NAME == "" {
		DB_NAME = "urlshortner"
	}

	DB_CONFIG := fmt.Sprintf("user=%v password=%v dbname=%v port=5432 sslmode=disable TimeZone=Asia/Bangkok", DB_USER, DB_PASS, DB_NAME)

	for {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  DB_CONFIG,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})

		if err == nil {
			break
		}

		fmt.Println("Failed to connect to database retry in 5 seconds")
		time.Sleep(time.Second * 5)
	}
}

func MigrationDB() {
	db.AutoMigrate(&dto.URLShortener{})
	db.AutoMigrate(&dto.User{})
}

func GetDB() *gorm.DB {
	return db
}
