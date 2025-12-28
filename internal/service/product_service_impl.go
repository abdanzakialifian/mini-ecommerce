package service

import (
	"context"
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

func (p productServiceImpl) GetProducts(ctx context.Context) ([]model.Product, *helper.AppError) {
	products, err := p.repository.FindAll(ctx)
	if err != nil {
		return nil, &helper.AppError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
			Err:        err,
		}
	}

	if len(products) == 0 {
		return nil, &helper.AppError{
			StatusCode: http.StatusNotFound,
			Message:    "Products Not Found",
			Err:        err,
		}
	}

	return products, nil
}
