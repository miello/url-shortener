package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/miello/url-shortener/backend/src/dto"
	"github.com/miello/url-shortener/backend/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateNewURL(ctx *gin.Context) {
	var body *dto.CreateURLDTO
	err := ctx.BindJSON(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing URL",
		})
		return
	}

	BASE_URL := os.Getenv("SHORTEN_BASE_URL")
	id := utils.GenerateNewId(5, body.Url)

	ctx.JSON(http.StatusCreated, gin.H{
		"url": fmt.Sprintf("%v/api/short/%v", BASE_URL, id),
	})
}

func RedirectToUrl(ctx *gin.Context) {
	id := ctx.Param("id")
	url, ok := utils.GetUrlFromId(id)

	if !ok {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Redirect(301, url)
}
