package service

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miello/url-shortener/backend/src/dto"
	"github.com/miello/url-shortener/backend/src/types"
	"github.com/miello/url-shortener/backend/src/utils"
	"golang.org/x/crypto/bcrypt"
)

func encryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalln("Something wrong with bcrypt hashing")
	}

	return string(hash)
}

func comparePassword(original string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(original), []byte(password))

	return err == nil
}

func Login(ctx *gin.Context) {
	var user types.LoginRequest
	db := utils.GetDB()
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing Username or Password",
		})
		return
	}

	var user_match dto.User

	username := user.Username
	db.Where(&dto.User{User: username}).Find(&user_match)

	if user_match.User == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Username and Password not correct",
		})
		return
	}

	password := user.Password
	isCorrect := comparePassword(user_match.Password, password)

	if isCorrect {
		token := GenerateToken(username, true)
		ctx.SetCookie("token", token, int(time.Hour.Minutes()), "/", "localhost", false, true)
		ctx.JSON(http.StatusAccepted, gin.H{
			"status": "success",
			"token":  token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Username and Password not correct",
		})
	}
}

func Register(ctx *gin.Context) {
	var user types.RegisterRequest
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing Username or Password or Handle",
		})
		return
	}

	db := utils.GetDB()
	username := user.Username

	var find_user dto.User

	db.Where(&dto.User{User: username}).First(&find_user)

	if find_user.User != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "This username is already used",
		})
		return
	}

	password := encryptPassword(user.Password)
	handle := user.Handle

	db.Create(&dto.User{
		User:     username,
		Password: password,
		Handle:   handle,
	})

	ctx.JSON(http.StatusCreated, "Register Successful")
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", 0, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, "Logout Successfully")
}
