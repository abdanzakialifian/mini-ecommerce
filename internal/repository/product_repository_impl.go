package repository

import (
	"context"
	"fmt"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewProductRepositoryImpl(db *pgxpool.Pool) domain.ProductRepository {
	return &productRepositoryImpl{db: db}
}

func (p *productRepositoryImpl) FindAll(ctx context.Context) ([]model.Product, error) {
	rows, err := p.db.Query(ctx, `SELECT id, category_id, name, description, price, stock, created_at, updated_at FROM products`)
	if err != nil {
		return nil, fmt.Errorf("Query product : %w", err)
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.CategoryID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Scan product %w", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Rows error : %w", err)
	}

	return products, nil
}
