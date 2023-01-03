package middlewares

import (
	"net/http"

	"project-go/handler"
	"project-go/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, &handler.Response{
				Success: false,
				Message: "Error: Something went wrong",
				Data:    "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
