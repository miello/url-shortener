package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/miello/url-shortener/backend/src/middleware"
	"github.com/miello/url-shortener/backend/src/service"
	"github.com/miello/url-shortener/backend/src/utils"
)

func main() {
	err := godotenv.Load()
	rand.Seed(time.Now().UnixNano())

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.ConnectDB()
	utils.MigrationDB()

	serve := gin.Default()
	serve.LoadHTMLGlob("templates/*")

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:5000"}

	serve.Use(cors.New(config))
	serve.Use(gin.Logger())

	url_api := serve.Group("/api/short")
	auth_api := serve.Group("/api/auth")

	{
		url_api.POST("", middleware.AuthorizeJWT(false), service.CreateNewURL)
		url_api.PUT("/:id", middleware.AuthorizeJWT(true), middleware.CheckAuthorMiddleware(), service.EditOriginalUrl)
		url_api.DELETE("/:id", middleware.AuthorizeJWT(true), middleware.CheckAuthorMiddleware(), service.DeleteShortenerUrl)
		serve.GET("/s/:id", service.RedirectToUrl)
		url_api.GET("/history", middleware.AuthorizeJWT(true), service.GetHistoryWithLimit)
	}

	{
		auth_api.POST("/login", service.Login)
		auth_api.POST("/register", service.Register)
		auth_api.PATCH("/logout", service.Logout)
		auth_api.PUT("/forgetpassword", service.ForgetPassword)
		auth_api.GET("/user", middleware.AuthorizeJWT(true), service.GetUser)
	}

	serve.Run()
}
