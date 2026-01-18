package service

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/category"
	"mini-ecommerce/internal/helper"
	"net/http"
)

type categoryServiceImpl struct {
	repository category.CategoryRepository
}

func NewCategoryServiceImpl(repository category.CategoryRepository) category.CategoryService {
	return &categoryServiceImpl{repository: repository}
}

func (c *categoryServiceImpl) CreateCategory(ctx context.Context, category *category.Category) *helper.AppError {
	err := c.repository.Create(ctx, category)
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

func (c *categoryServiceImpl) GetCategory(ctx context.Context, id string) (category.Category, *helper.AppError) {
	result, err := c.repository.Find(ctx, id)
	if err != nil {
		if errors.Is(err, helper.ErrCategoryNotFound) {
			return category.Category{}, helper.NewAppError(
				http.StatusNotFound,
				"Product Not Found",
				err,
			)
		}

		return category.Category{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return result, nil
}

func (c *categoryServiceImpl) GetCategories(ctx context.Context) ([]category.Category, *helper.AppError) {
	results, err := c.repository.FindAll(ctx)
	if err != nil {
		return nil, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return results, nil
}

func (c *categoryServiceImpl) UpdateCategory(ctx context.Context, updateCategory *category.UpdateCategory) *helper.AppError {
	err := c.repository.Update(ctx, updateCategory)
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

func (c *categoryServiceImpl) DeleteCategory(ctx context.Context, id string) *helper.AppError {
	err := c.repository.Delete(ctx, id)
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
