package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miello/url-shortener/backend/src/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken, err := ctx.Cookie("token")
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		_, validateErr := service.ValidateToken(authToken)
		if validateErr != nil {
			fmt.Println(validateErr)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
