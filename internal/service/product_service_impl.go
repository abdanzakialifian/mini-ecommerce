package service

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"
	"net/http"
)

type productServiceImpl struct {
	repository domain.ProductRepository
}

func NewProductServiceImpl(repository domain.ProductRepository) domain.ProductService {
	return &productServiceImpl{repository: repository}
}

func (p *productServiceImpl) CreateProduct(ctx context.Context, product *model.Product) *helper.AppError {
	err := p.repository.Create(ctx, product)
	if err != nil {
		if errors.Is(err, domain.ErrProductAlreadyExists) {
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

func (p *productServiceImpl) GetProduct(ctx context.Context, id string) (model.Product, *helper.AppError) {
	product, err := p.repository.Find(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return model.Product{}, helper.NewAppError(
				http.StatusNotFound,
				"Product Not Found",
				err,
			)
		}

		return model.Product{}, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return product, nil
}

func (p productServiceImpl) GetProducts(ctx context.Context) ([]model.Product, *helper.AppError) {
	products, err := p.repository.FindAll(ctx)
	if err != nil {
		return nil, helper.NewAppError(
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
	}

	return products, nil
}

func (p *productServiceImpl) UpdateProduct(ctx context.Context, product *model.UpdateProduct) *helper.AppError {
	err := p.repository.Update(ctx, product)

	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
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
		if errors.Is(err, domain.ErrProductNotFound) {
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
