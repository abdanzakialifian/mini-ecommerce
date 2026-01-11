package handler

import (
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/handler/request"
	"mini-ecommerce/internal/handler/response"
	"mini-ecommerce/internal/helper"
	"net/http"
	"strconv"

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

	user := toUserFromCreate(req)
	appErr := h.service.CreateUser(c.Request.Context(), &user)

	if appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Create User",
		toUserResponse(user),
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

	loginUser := toLoginUser(req)
	user, accessToken, appErr := h.service.GetUserByEmail(c.Request.Context(), loginUser)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Get User",
		toLoginUserResponse(user, accessToken),
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

	updateCategory := toUserFromUpdate(req, userId)

	appErr := h.service.UpdateUser(c.Request.Context(), &updateCategory)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Update User",
		toUpdateUserResponse(updateCategory),
	)
	c.JSON(status, res)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("User id is required"),
		))
		return
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			errors.New("Failed convert id"),
		))
		return
	}

	appErr := h.service.DeleteUser(c.Request.Context(), newId)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Update User")
	c.JSON(status, res)
}
