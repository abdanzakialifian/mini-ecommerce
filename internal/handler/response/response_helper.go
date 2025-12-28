package response

import (
	"net/http"
)

func Success(message string, data any) (int, BaseResponse) {
	return http.StatusOK, BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func Created(message string, data any) (int, BaseResponse) {
	return http.StatusCreated, BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func Error(message string, err any, status int) (int, BaseResponse) {
	return status, BaseResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
