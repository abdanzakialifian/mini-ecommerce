package repository

import (
	"context"
	"mini-ecommerce/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewProductRepositoryImpl(db *pgxpool.Pool) ProductRepository {
	return &productRepositoryImpl{db: db}
}

func (p *productRepositoryImpl) FindAll() ([]model.Product, error) {
	rows, err := p.db.Query(context.Background(), `SELECT * FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.CategoryID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
