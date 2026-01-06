package service

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"
	"net/http"
)

type categoryServiceImpl struct {
	repository domain.CategoryRepository
}

func NewCategoryServiceImpl(repository domain.CategoryRepository) domain.CategoryService {
	return &categoryServiceImpl{repository: repository}
}

func (c *categoryServiceImpl) CreateCategory(ctx context.Context, category *model.Category) *helper.AppError {
	err := c.repository.Create(ctx, category)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryAlreadyExists) {
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

func (c *categoryServiceImpl) GetCategory(ctx context.Context, id string) (model.Category, *helper.AppError) {
	category, err := c.repository.Find(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			return model.Category{}, helper.NewAppError(
				http.StatusNotFound,
				"Product Not Found",
				err,
			)
		}

		return model.Category{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return category, nil
}

func (c *categoryServiceImpl) GetCategories(ctx context.Context) ([]model.Category, *helper.AppError) {
	categories, err := c.repository.FindAll(ctx)
	if err != nil {
		return nil, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return categories, nil
}

func (c *categoryServiceImpl) UpdateCategory(ctx context.Context, category *model.Category) *helper.AppError {
	err := c.repository.Update(ctx, category)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
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

func (c *categoryServiceImpl) DeleteCategory(ctx context.Context, id string) *helper.AppError {
	err := c.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
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
