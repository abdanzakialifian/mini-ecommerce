package handler

import (
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/handler/request"
	"mini-ecommerce/internal/handler/response"
	"mini-ecommerce/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if appErr := h.service.CreateUser(c.Request.Context(), &user); appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := response.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	status, res := response.Success(
		"Success Create User",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	var req request.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	loginUser := model.LoginUser{
		Email:    req.Email,
		Password: req.Password,
	}

	user, accessToken, appErr := h.service.GetUserByEmail(c.Request.Context(), loginUser)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := response.LoginUserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: accessToken,
	}

	status, res := response.Success(
		"Success Get User",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userId := c.MustGet("user_id").(int)
	var req request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	updateUser := model.UpdateUser{
		ID:          userId,
		Name:        req.Name,
		Email:       req.Email,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}

	if appErr := h.service.UpdateUser(c.Request.Context(), &updateUser); appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := response.UserResponse{
		ID:    updateUser.ID,
		Name:  *updateUser.Name,
		Email: *updateUser.Email,
	}

	status, res := response.Success(
		"Success Update User",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	if appErr := h.service.DeleteUser(c.Request.Context(), userId); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Delete User")
	c.JSON(status, res)
}
