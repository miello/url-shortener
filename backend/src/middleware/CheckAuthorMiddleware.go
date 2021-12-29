package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/miello/url-shortener/backend/src/dto"
	"github.com/miello/url-shortener/backend/src/utils"
)

func CheckAuthorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Missing Shortener ID",
			})
			return
		}

		var find_user dto.URLShortener

		db := utils.GetDB()
		find_err := db.Where(&dto.URLShortener{UrlID: id}).First(&find_user)
		if find_err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not found this id",
			})
			return
		}

		userId, _ := ctx.Get("user")

		var uid uuid.UUID
		var uuid_err error

		uid, uuid_err = uuid.Parse(userId.(string))
		if uuid_err != nil || uid.String() != find_user.UserID.String() {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not found this id",
			})
			return
		}

		ctx.Set("post", find_user)
		ctx.Next()
	}
}
