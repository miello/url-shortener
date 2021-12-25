package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/miello/url-shortener/backend/src/dto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB_USER := os.Getenv("DATABASE_USER")
	DB_PASS := os.Getenv("DATABASE_PASSWORD")
	DB_NAME := os.Getenv("DATABASE_NAME")

	DB_CONFIG := fmt.Sprintf("user=%v password=%v dbname=%v port=5432 sslmode=disable TimeZone=Asia/Bangkok", DB_USER, DB_PASS, DB_NAME)

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  DB_CONFIG,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB.AutoMigrate(&dto.URLShortenerDTO{})
	DB.AutoMigrate(&dto.UserDTO{})

	DB.Create(&dto.UserDTO{
		Handle:   "earthza",
		User:     "partearth",
		Password: "earthza1412",
	})
}
