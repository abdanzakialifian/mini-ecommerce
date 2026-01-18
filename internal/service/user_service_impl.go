package service

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/user"
	"mini-ecommerce/internal/helper"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repository user.UserRepository
}

func NewUserServiceImpl(repository user.UserRepository) user.UserService {
	return &userServiceImpl{repository: repository}
}

func (u *userServiceImpl) CreateUser(ctx context.Context, user *user.User) *helper.AppError {
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
		if errors.Is(err, helper.ErrUserAlreadyExists) {
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

func (u *userServiceImpl) GetUserByEmail(ctx context.Context, login user.LoginUser) (user.User, string, *helper.AppError) {
	result, accessToken, err := u.repository.FindByEmail(ctx, login)
	if err != nil {
		if errors.Is(err, helper.ErrUserInvalid) {
			return user.User{}, "", helper.NewAppError(
				http.StatusBadRequest,
				"Validation Failed",
				err,
			)
		}

		return user.User{}, "", helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return result, accessToken, nil
}

func (u *userServiceImpl) UpdateUser(ctx context.Context, updateUser *user.UpdateUser) *helper.AppError {
	if updateUser.OldPassword != nil || updateUser.NewPassword != nil {
		if updateUser.OldPassword == nil || updateUser.NewPassword == nil {
			return helper.NewAppError(
				http.StatusBadRequest,
				"Invalid Request",
				errors.New("Both old password and new password are required"),
			)
		}

		result, err := u.repository.FindById(ctx, updateUser.ID)
		if err != nil {
			if errors.Is(err, helper.ErrUserNotFound) {
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

		if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(*updateUser.OldPassword)); err != nil {
			return helper.NewAppError(
				http.StatusBadRequest,
				"Invalid Request",
				errors.New("Old password is incorrect"),
			)
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(*updateUser.NewPassword), bcrypt.DefaultCost)

		if err != nil {
			return helper.NewAppError(
				http.StatusInternalServerError,
				"Internal Server Error",
				err,
			)
		}

		newPasswordHash := string(hash)
		updateUser.NewPassword = &newPasswordHash
	}

	err := u.repository.Update(ctx, updateUser)
	if err != nil {
		if errors.Is(err, helper.ErrUserNotFound) {
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

func (u *userServiceImpl) DeleteUser(ctx context.Context, id int) *helper.AppError {
	err := u.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, helper.ErrUserNotFound) {
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
