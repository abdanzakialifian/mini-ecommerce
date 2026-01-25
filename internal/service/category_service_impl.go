package service

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/category"
	"mini-ecommerce/internal/helper"
	"net/http"
)

type categoryServiceImpl struct {
	categoryRepository category.Repository
}

func NewCategory(categoryRepository category.Repository) category.Service {
	return &categoryServiceImpl{categoryRepository: categoryRepository}
}

func (c *categoryServiceImpl) Create(ctx context.Context, data *category.Data) *helper.AppError {
	err := c.categoryRepository.Create(ctx, data)
	if err != nil {
		if errors.Is(err, helper.ErrCategoryAlreadyExists) {
			return helper.NewAppError(
				http.StatusConflict,
				"Product Already Exists",
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

func (c *categoryServiceImpl) Get(ctx context.Context, id string) (category.Data, *helper.AppError) {
	categoryData, err := c.categoryRepository.Find(ctx, id)
	if err != nil {
		if errors.Is(err, helper.ErrCategoryNotFound) {
			return category.Data{}, helper.NewAppError(
				http.StatusNotFound,
				"Product Not Found",
				err,
			)
		}

		return category.Data{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return categoryData, nil
}

func (c *categoryServiceImpl) GetAll(ctx context.Context) ([]category.Data, *helper.AppError) {
	categories, err := c.categoryRepository.FindAll(ctx)
	if err != nil {
		return nil, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return categories, nil
}

func (c *categoryServiceImpl) Update(ctx context.Context, update *category.Update) *helper.AppError {
	err := c.categoryRepository.Update(ctx, update)
	if err != nil {
		if errors.Is(err, helper.ErrCategoryNotFound) {
			return helper.NewAppError(
				http.StatusNotFound,
				"Product Not Found",
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

func (c *categoryServiceImpl) Delete(ctx context.Context, id string) *helper.AppError {
	err := c.categoryRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, helper.ErrCategoryNotFound) {
			return helper.NewAppError(
				http.StatusNotFound,
				"Product Not Found",
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
