package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/miello/url-shortener/backend/src/service"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serve := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5000"}

	serve.Use(cors.Default())

	url_api := serve.Group("/api/short")
	url_api.Use(cors.Default())

	{
		url_api.POST("", service.CreateNewURL)
	}

	serve.GET("/s/:id", service.RedirectToUrl)

	serve.Run()
}
