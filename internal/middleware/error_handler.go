package middleware

import (
	"errors"
	"mini-ecommerce/internal/handler/response"
	"mini-ecommerce/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		var appErr *helper.AppError

		if errors.As(err, &appErr) {
			status, res := response.Error(
				appErr.Message,
				func() any {
					if appErr.Err != nil {
						return appErr.Err.Error()
					}
					return nil
				}(),
				appErr.StatusCode,
			)
			c.JSON(status, res)
			return
		}

		status, res := response.Error(
			"Internal Server Error",
			err.Error(),
			http.StatusInternalServerError,
		)
		c.JSON(status, res)
	}
}
