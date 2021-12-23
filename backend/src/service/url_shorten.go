package service

import (
	"net/http"

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

	id := utils.GenerateNewId(5, body.Url)

	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
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
