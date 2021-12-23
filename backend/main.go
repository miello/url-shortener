package main

import (
	"math/rand"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/miello/url-shortener/backend/src/service"
)

func main() {
	rand.Seed(time.Hour.Nanoseconds())

	serve := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5000"}

	serve.Use(cors.Default())

	url_api := serve.Group("/api/short")
	url_api.Use(cors.Default())

	{
		url_api.POST("", service.CreateNewURL)
		url_api.GET("/:id", service.RedirectToUrl)
	}

	serve.Run()
}
