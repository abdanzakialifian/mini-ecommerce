package service

import (
	"mini-ecommerce/internal/model"
)

type ProductService interface {
	GetProducts() ([]model.Product, error)
}
