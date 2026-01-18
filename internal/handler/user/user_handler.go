package user

import (
	"mini-ecommerce/internal/domain/user"
	"mini-ecommerce/internal/helper"
	"mini-ecommerce/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service user.UserService
}

func NewHandler(service user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	user := user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if appErr := h.service.CreateUser(c.Request.Context(), &user); appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := UserResponse{
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
	var req LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	loginUser := user.LoginUser{
		Email:    req.Email,
		Password: req.Password,
	}

	result, accessToken, appErr := h.service.GetUserByEmail(c.Request.Context(), loginUser)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := LoginUserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
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

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	updateUser := user.UpdateUser{
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

	dataResponse := UserResponse{
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
