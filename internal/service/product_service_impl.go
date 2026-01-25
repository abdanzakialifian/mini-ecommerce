package service

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/product"
	"mini-ecommerce/internal/helper"
	"net/http"
)

type productServiceImpl struct {
	productRepository product.Repository
}

func NewProduct(productRepository product.Repository) product.Service {
	return &productServiceImpl{productRepository: productRepository}
}

func (p *productServiceImpl) Create(ctx context.Context, data *product.Data) *helper.AppError {
	err := p.productRepository.Create(ctx, data)
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

func (p *productServiceImpl) Get(ctx context.Context, id string) (product.Data, *helper.AppError) {
	productData, err := p.productRepository.Find(ctx, id)
	if err != nil {
		if errors.Is(err, helper.ErrProductNotFound) {
			return product.Data{}, helper.NewAppError(
				http.StatusNotFound,
				"Product Not Found",
				err,
			)
		}

		return product.Data{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return productData, nil
}

func (p productServiceImpl) GetAll(ctx context.Context) ([]product.Data, *helper.AppError) {
	products, err := p.productRepository.FindAll(ctx)
	if err != nil {
		return nil, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return products, nil
}

func (p *productServiceImpl) Update(ctx context.Context, update *product.Update) *helper.AppError {
	err := p.productRepository.Update(ctx, update)

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

func (p *productServiceImpl) Delete(ctx context.Context, id string) *helper.AppError {
	err := p.productRepository.Delete(ctx, id)

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
