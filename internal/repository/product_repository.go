package repository

import (
	"mini-ecommerce/internal/model"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
}
