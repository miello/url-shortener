package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miello/url-shortener/backend/src/service"
)

func AuthorizeJWT(strict bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken, err := ctx.Cookie("access_token")
		if err != nil && strict {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Please Login First",
			})
			return
		} else if err != nil && !strict {
			ctx.Next()
			return
		}

		token, validateErr := service.ValidateToken(authToken)

		if validateErr != nil || !token.Valid {
			fmt.Println(validateErr)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Please try to logout and login",
			})
			return
		} else if !strict && validateErr != nil {
			ctx.Next()
			return
		}

		claims, ok := service.ParseToken(token)

		if ok {
			ctx.Set("user", claims.UserID)
			ctx.Next()
		} else if strict {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Please try to logout and login",
			})
		} else if !strict {
			ctx.Next()
		}

	}
}
