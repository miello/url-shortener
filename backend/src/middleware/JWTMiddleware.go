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
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		} else if err != nil && !strict {
			ctx.Next()
			return
		}

		token, validateErr := service.ValidateToken(authToken)

		if validateErr != nil || !token.Valid {
			fmt.Println(validateErr)
			ctx.AbortWithStatus(http.StatusUnauthorized)
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
			ctx.AbortWithStatus(http.StatusUnauthorized)
		} else if !strict {
			ctx.Next()
		}

	}
}
