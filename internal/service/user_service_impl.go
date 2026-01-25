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
	userRepository user.Repository
}

func NewUser(userRepository user.Repository) user.Service {
	return &userServiceImpl{userRepository: userRepository}
}

func (u *userServiceImpl) Create(ctx context.Context, data *user.Data) *helper.AppError {
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	data.Password = string(hash)

	err = u.userRepository.Create(ctx, data)

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

func (u *userServiceImpl) GetByEmail(ctx context.Context, login user.Login) (user.Data, string, *helper.AppError) {
	userData, accessToken, err := u.userRepository.FindByEmail(ctx, login)
	if err != nil {
		if errors.Is(err, helper.ErrUserInvalid) {
			return user.Data{}, "", helper.NewAppError(
				http.StatusBadRequest,
				"Validation Failed",
				err,
			)
		}

		return user.Data{}, "", helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return userData, accessToken, nil
}

func (u *userServiceImpl) Update(ctx context.Context, update *user.Update) *helper.AppError {
	if update.OldPassword != nil || update.NewPassword != nil {
		if update.OldPassword == nil || update.NewPassword == nil {
			return helper.NewAppError(
				http.StatusBadRequest,
				"Invalid Request",
				errors.New("Both old password and new password are required"),
			)
		}

		userData, err := u.userRepository.FindById(ctx, update.ID)
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

		if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(*update.OldPassword)); err != nil {
			return helper.NewAppError(
				http.StatusBadRequest,
				"Invalid Request",
				errors.New("Old password is incorrect"),
			)
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(*update.NewPassword), bcrypt.DefaultCost)

		if err != nil {
			return helper.NewAppError(
				http.StatusInternalServerError,
				"Internal Server Error",
				err,
			)
		}

		newPasswordHash := string(hash)
		update.NewPassword = &newPasswordHash
	}

	err := u.userRepository.Update(ctx, update)
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

func (u *userServiceImpl) Delete(ctx context.Context, id int) *helper.AppError {
	err := u.userRepository.Delete(ctx, id)
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
