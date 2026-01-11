package handler

import (
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/handler/request"
	"mini-ecommerce/internal/handler/response"
)

func toUserFromCreate(req request.CreateUserRequest) model.User {
	return model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func toUserFromUpdate(req request.UpdateUserRequest) model.UpdateUser {
	return model.UpdateUser{
		ID:       req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func toUserResponse(user model.User) response.UserResponse {
	return response.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func toUpdateUserResponse(user model.UpdateUser) response.UserResponse {
	return response.UserResponse{
		ID:       user.ID,
		Name:     *user.Name,
		Email:    *user.Email,
		Password: *user.Password,
	}
}

func toLoginUser(req request.LoginUserRequest) model.LoginUser {
	return model.LoginUser{
		Email:    req.Email,
		Password: req.Password,
	}
}

func toLoginUserResponse(user model.User, accessToken string) response.LoginUserResponse {
	return response.LoginUserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: accessToken,
	}
}
