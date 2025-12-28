package service

import (
	"mini-ecommerce/internal/model"
	"mini-ecommerce/internal/repository"
)

type productServiceImpl struct {
	repository repository.ProductRepository
}

func NewProductServiceImpl(repository repository.ProductRepository) ProductService {
	return &productServiceImpl{repository: repository}
}

func (p productServiceImpl) GetProducts() ([]model.Product, error) {
	return p.repository.FindAll()
}
