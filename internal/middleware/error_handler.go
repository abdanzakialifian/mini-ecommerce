package middleware

import (
	"errors"
	"mini-ecommerce/internal/helper"
	"mini-ecommerce/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		for _, ginErr := range c.Errors {
			var appErr *helper.AppError

			if errors.As(ginErr, &appErr) {
				var detail any
				if appErr.Err != nil {
					detail = appErr.Err.Error()
				}

				status, res := response.Error(
					appErr.Message,
					detail,
					appErr.StatusCode,
				)

				c.JSON(status, res)
				c.Abort()
				return
			}
		}

		lastErr := c.Errors.Last().Err

		status, res := response.Error(
			"Internal Server Error",
			lastErr.Error(),
			http.StatusInternalServerError,
		)

		c.JSON(status, res)
		c.Abort()
	}
}
