package service

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/product"
	"mini-ecommerce/internal/helper"
	"net/http"
)

type productServiceImpl struct {
	repository product.ProductRepository
}

func NewProductServiceImpl(repository product.ProductRepository) product.ProductService {
	return &productServiceImpl{repository: repository}
}

func (p *productServiceImpl) CreateProduct(ctx context.Context, product *product.Product) *helper.AppError {
	err := p.repository.Create(ctx, product)
	if err != nil {
		if errors.Is(err, helper.ErrProductAlreadyExists) {
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

func (p *productServiceImpl) GetProduct(ctx context.Context, id string) (product.Product, *helper.AppError) {
	result, err := p.repository.Find(ctx, id)
	if err != nil {
		if errors.Is(err, helper.ErrProductNotFound) {
			return product.Product{}, helper.NewAppError(
				http.StatusNotFound,
				"Product Not Found",
				err,
			)
		}

		return product.Product{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return result, nil
}

func (p productServiceImpl) GetProducts(ctx context.Context) ([]product.Product, *helper.AppError) {
	results, err := p.repository.FindAll(ctx)
	if err != nil {
		return nil, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return results, nil
}

func (p *productServiceImpl) UpdateProduct(ctx context.Context, updateProduct *product.UpdateProduct) *helper.AppError {
	err := p.repository.Update(ctx, updateProduct)

	if err != nil {
		if errors.Is(err, helper.ErrProductNotFound) {
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

func (p *productServiceImpl) DeleteProduct(ctx context.Context, id string) *helper.AppError {
	err := p.repository.Delete(ctx, id)

	if err != nil {
		if errors.Is(err, helper.ErrProductNotFound) {
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
