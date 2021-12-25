package service

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/miello/url-shortener/backend/src/types"
	"github.com/miello/url-shortener/backend/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateNewURL(ctx *gin.Context) {
	var body *types.CreateURLRequest
	err := ctx.BindJSON(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing URL",
		})
		return
	}

	containHttp := strings.Contains(body.Url, "http://")
	containHttps := strings.Contains(body.Url, "https://")

	if !containHttp && !containHttps {
		body.Url = "http://" + body.Url
	}

	BASE_URL := os.Getenv("SHORTEN_BASE_URL")
	id := utils.GenerateNewId(5, body.Url)

	ctx.JSON(http.StatusCreated, gin.H{
		"url": fmt.Sprintf("%v/s/%v", BASE_URL, id),
	})
}

func RedirectToUrl(ctx *gin.Context) {
	id := ctx.Param("id")
	url, ok := utils.GetUrlFromId(id)

	if !ok {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url)
}
