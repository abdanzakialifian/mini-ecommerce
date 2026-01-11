package middleware

import (
	"errors"
	"mini-ecommerce/internal/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(helper.NewAppError(
				http.StatusUnauthorized,
				"Authorization token is required",
				errors.New("Missing authorization header"),
			))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Error(helper.NewAppError(
				http.StatusUnauthorized,
				"Authorization token is required",
				errors.New("Invalid authorization format"),
			))
			c.Abort()
			return
		}

		accessToken := parts[1]
		claims, err := helper.ParseToken(accessToken)

		if err != nil {
			c.Error(helper.NewAppError(
				http.StatusUnauthorized,
				"Authorization token is required",
				errors.New("Invalid or expired token"),
			))
			c.Abort()
			return
		}

		id := claims["id"].(float64)
		c.Set("user_id", int(id))
		c.Next()
	}
}
