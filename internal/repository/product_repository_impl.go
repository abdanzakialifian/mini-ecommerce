package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/product"
	"mini-ecommerce/internal/helper"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepositoryImpl struct {
	db *pgxpool.Pool
	tx *helper.Transaction
}

func NewProductRepositoryImpl(db *pgxpool.Pool, tx *helper.Transaction) product.ProductRepository {
	return &productRepositoryImpl{db: db, tx: tx}
}

func (p *productRepositoryImpl) Create(ctx context.Context, product *product.Product) error {
	query := "INSERT INTO products (category_id, name, description, price, stock) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := p.db.QueryRow(
		ctx,
		query,
		product.CategoryID,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
	).Scan(&product.ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return helper.ErrProductAlreadyExists
		}
		return err
	}

	return nil
}

func (p *productRepositoryImpl) Find(ctx context.Context, id string) (product.Product, error) {
	query := "SELECT id, category_id, name, description, price, stock FROM products WHERE id = $1"
	var result product.Product
	err := p.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&result.ID,
		&result.CategoryID,
		&result.Name,
		&result.Description,
		&result.Price,
		&result.Stock,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return product.Product{}, helper.ErrProductNotFound
		}

		return product.Product{}, err
	}

	return result, nil
}

func (p *productRepositoryImpl) FindAll(ctx context.Context) ([]product.Product, error) {
	query := "SELECT id, category_id, name, description, price, stock FROM products"
	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []product.Product

	for rows.Next() {
		var product product.Product
		if err := rows.Scan(
			&product.ID,
			&product.CategoryID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
		); err != nil {
			return nil, err
		}
		results = append(results, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (p *productRepositoryImpl) Update(ctx context.Context, updateProduct *product.UpdateProduct) error {
	query := "UPDATE products SET category_id = COALESCE($1, category_id), name = COALESCE($2, name), description = COALESCE($3, description), price = COALESCE($4, price), stock = COALESCE($5, stock), updated_at = NOW() WHERE id = $6 RETURNING id, category_id, name, description, price, stock"
	err := p.db.QueryRow(
		ctx,
		query,
		updateProduct.CategoryID,
		updateProduct.Name,
		updateProduct.Description,
		updateProduct.Price,
		updateProduct.Stock,
		updateProduct.ID,
	).Scan(
		&updateProduct.ID,
		&updateProduct.CategoryID,
		&updateProduct.Name,
		&updateProduct.Description,
		&updateProduct.Price,
		&updateProduct.Stock,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return helper.ErrProductNotFound
		}
		return err
	}

	return nil
}

func (p *productRepositoryImpl) UpdateStock(ctx context.Context, id string, quantity int) error {
	db := p.tx.GetTx(ctx)
	query := "UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1"
	cmd, err := db.Exec(ctx, query, quantity, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrProductInsufficientStock
	}

	return nil
}

func (p *productRepositoryImpl) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM products WHERE id = $1"
	cmd, err := p.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrProductNotFound
	}

	return nil
}
