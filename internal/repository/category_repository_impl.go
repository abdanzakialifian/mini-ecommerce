package repository

import (
	"context"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type categoryRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewCategoryRepositoryImpl(db *pgxpool.Pool) domain.CategoryRepository {
	return &categoryRepositoryImpl{db: db}
}

func (c *categoryRepositoryImpl) Create(ctx context.Context, category *model.Category) error {
	panic("unimplemented")
}

func (c *categoryRepositoryImpl) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (c *categoryRepositoryImpl) Find(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (c *categoryRepositoryImpl) FindAll(ctx context.Context) ([]model.Category, error) {
	panic("unimplemented")
}

func (c *categoryRepositoryImpl) Update(ctx context.Context, category *model.Category) error {
	panic("unimplemented")
}
