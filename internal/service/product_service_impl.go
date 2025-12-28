package service

import (
	"context"
	"fmt"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
)

type productServiceImpl struct {
	repository domain.ProductRepository
}

func NewProductServiceImpl(repository domain.ProductRepository) domain.ProductService {
	return &productServiceImpl{repository: repository}
}

func (p productServiceImpl) GetProducts(ctx context.Context) ([]model.Product, error) {
	products, err := p.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("Get products : %w", err)
	}
	return products, nil
}
