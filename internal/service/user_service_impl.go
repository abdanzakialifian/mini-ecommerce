package service

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repository domain.UserRepository
}

func NewUserServiceImpl(repository domain.UserRepository) domain.UserService {
	return userServiceImpl{repository: repository}
}

func (u userServiceImpl) CreateUser(ctx context.Context, user *model.User) *helper.AppError {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	user.Password = string(hash)

	err = u.repository.Create(ctx, user)

	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return helper.NewAppError(
				http.StatusConflict,
				"User Already Exists",
				err,
			)
		}

		return helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return nil
}

func (u userServiceImpl) GetUser(ctx context.Context, login model.LoginUser) (model.User, *helper.AppError) {
	user, err := u.repository.Find(ctx, login)
	if err != nil {
		if errors.Is(err, domain.ErrUserInvalid) {
			return model.User{}, helper.NewAppError(
				http.StatusBadRequest,
				"Validation Failed",
				err,
			)
		}

		return model.User{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return user, nil
}

func (u userServiceImpl) UpdateUser(ctx context.Context, updateUser *model.UpdateUser) *helper.AppError {
	err := u.repository.Update(ctx, updateUser)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return helper.NewAppError(
				http.StatusNotFound,
				"User Not Found",
				err,
			)
		}

		return helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return nil
}

func (u userServiceImpl) DeleteUser(ctx context.Context, id int) *helper.AppError {
	err := u.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return helper.NewAppError(
				http.StatusNotFound,
				"User Not Found",
				err,
			)
		}

		return helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return nil
}
