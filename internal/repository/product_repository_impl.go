package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewProductRepositoryImpl(db *pgxpool.Pool) domain.ProductRepository {
	return &productRepositoryImpl{db: db}
}

func (p *productRepositoryImpl) Create(ctx context.Context, product *model.Product) error {
	query := `INSERT INTO products (category_id, name, description, price, stock) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	err := p.db.QueryRow(
		ctx,
		query,
		product.CategoryID,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
	).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.ErrProductAlreadyExists
		}
		return err
	}

	return nil
}

func (p *productRepositoryImpl) FindAll(ctx context.Context) ([]model.Product, error) {
	rows, err := p.db.Query(ctx, `SELECT id, category_id, name, description, price, stock, created_at, updated_at FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(
			&product.ID,
			&product.CategoryID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
