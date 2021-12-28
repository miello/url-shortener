package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miello/url-shortener/backend/src/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		token, validateErr := service.ValidateToken(authToken)
		if validateErr != nil || !token.Valid {
			fmt.Println(validateErr)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		claims, ok := service.ParseToken(token)

		if ok {
			ctx.Set("user", claims.Name)
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
