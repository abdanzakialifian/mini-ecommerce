package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"

	"github.com/jackc/pgx/v5"
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
			return domain.ErrProductAlreadyExists
		}
		return err
	}

	return nil
}

func (p *productRepositoryImpl) Find(ctx context.Context, id string) (model.Product, error) {
	query := "SELECT id, category_id, name, description, price, stock FROM products WHERE id = $1"
	var product model.Product
	err := p.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&product.ID,
		&product.CategoryID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Product{}, domain.ErrProductNotFound
		}

		return model.Product{}, err
	}

	return product, nil
}

func (p *productRepositoryImpl) FindAll(ctx context.Context) ([]model.Product, error) {
	query := "SELECT id, category_id, name, description, price, stock FROM products"
	rows, err := p.db.Query(ctx, query)
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
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepositoryImpl) Update(ctx context.Context, product model.Product) (model.Product, error) {
	query := "UPDATE products SET category_id = COALESCE($1, category_id), name = COALESCE($2, name), description = COALESCE($3, description), price = COALESCE($4, price), stock = COALESCE($5, stock), updated_at = NOW() WHERE id = $6 RETURNING id, category_id, name, description, price, stock"
	updatedProduct := model.Product{}
	err := p.db.QueryRow(
		ctx,
		query,
		product.CategoryID,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.ID,
	).Scan(
		&updatedProduct.ID,
		&updatedProduct.CategoryID,
		&updatedProduct.Name,
		&updatedProduct.Description,
		&updatedProduct.Price,
		&updatedProduct.Stock,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Product{}, domain.ErrProductNotFound
		}
		return model.Product{}, err
	}

	return updatedProduct, nil
}

func (p *productRepositoryImpl) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM products WHERE id = $1"
	cmd, err := p.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return domain.ErrProductNotFound
	}

	return nil
}
