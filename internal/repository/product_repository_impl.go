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

func NewProduct(db *pgxpool.Pool, tx *helper.Transaction) product.Repository {
	return &productRepositoryImpl{db: db, tx: tx}
}

func (p *productRepositoryImpl) Create(ctx context.Context, data *product.Data) error {
	query := "INSERT INTO products (category_id, name, description, price, stock) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := p.db.QueryRow(
		ctx,
		query,
		data.CategoryID,
		data.Name,
		data.Description,
		data.Price,
		data.Stock,
	).Scan(&data.ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return helper.ErrProductAlreadyExists
		}
		return err
	}

	return nil
}

func (p *productRepositoryImpl) Find(ctx context.Context, id string) (product.Data, error) {
	query := "SELECT id, category_id, name, description, price, stock FROM products WHERE id = $1"
	var productData product.Data
	err := p.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&productData.ID,
		&productData.CategoryID,
		&productData.Name,
		&productData.Description,
		&productData.Price,
		&productData.Stock,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return product.Data{}, helper.ErrProductNotFound
		}

		return product.Data{}, err
	}

	return productData, nil
}

func (p *productRepositoryImpl) FindAll(ctx context.Context) ([]product.Data, error) {
	query := "SELECT id, category_id, name, description, price, stock FROM products"
	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []product.Data

	for rows.Next() {
		var productData product.Data
		if err := rows.Scan(
			&productData.ID,
			&productData.CategoryID,
			&productData.Name,
			&productData.Description,
			&productData.Price,
			&productData.Stock,
		); err != nil {
			return nil, err
		}
		products = append(products, productData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepositoryImpl) Update(ctx context.Context, update *product.Update) error {
	query := "UPDATE products SET category_id = COALESCE($1, category_id), name = COALESCE($2, name), description = COALESCE($3, description), price = COALESCE($4, price), stock = COALESCE($5, stock), updated_at = NOW() WHERE id = $6 RETURNING id, category_id, name, description, price, stock"
	err := p.db.QueryRow(
		ctx,
		query,
		update.CategoryID,
		update.Name,
		update.Description,
		update.Price,
		update.Stock,
		update.ID,
	).Scan(
		&update.ID,
		&update.CategoryID,
		&update.Name,
		&update.Description,
		&update.Price,
		&update.Stock,
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
