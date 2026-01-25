package user

import (
	"mini-ecommerce/internal/domain/user"
	"mini-ecommerce/internal/helper"
	"mini-ecommerce/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.Service
}

func NewHandler(userService user.Service) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	userData := user.Data{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if appErr := h.userService.Create(c.Request.Context(), &userData); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Create User",
		Response{
			ID:    userData.ID,
			Name:  userData.Name,
			Email: userData.Email,
		},
	)
	c.JSON(status, res)
}

func (h *UserHandler) GetByEmail(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	login := user.Login{
		Email:    req.Email,
		Password: req.Password,
	}
	result, accessToken, appErr := h.userService.GetByEmail(c.Request.Context(), login)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Get User",
		LoginResponse{
			ID:          result.ID,
			Name:        result.Name,
			Email:       result.Email,
			AccessToken: accessToken,
		},
	)
	c.JSON(status, res)
}

func (h *UserHandler) Update(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	userUpdate := user.Update{
		ID:          userId,
		Name:        req.Name,
		Email:       req.Email,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
	if appErr := h.userService.Update(c.Request.Context(), &userUpdate); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Update User",
		Response{
			ID:    userUpdate.ID,
			Name:  *userUpdate.Name,
			Email: *userUpdate.Email,
		},
	)
	c.JSON(status, res)
}

func (h *UserHandler) Delete(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	if appErr := h.userService.Delete(c.Request.Context(), userId); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Delete User")
	c.JSON(status, res)
}
