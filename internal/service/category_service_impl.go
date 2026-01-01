package service

import (
	"context"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
)

type categoryServiceImpl struct {
	repository domain.CategoryRepository
}

func NewCategoryServiceImpl(repository domain.CategoryRepository) domain.CategoryService {
	return &categoryServiceImpl{repository: repository}
}

func (c *categoryServiceImpl) CreateCategory(ctx context.Context, category *model.Category) error {
	panic("unimplemented")
}

func (c *categoryServiceImpl) DeleteCategory(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (c *categoryServiceImpl) GetCategories(ctx context.Context) ([]model.Category, error) {
	panic("unimplemented")
}

func (c *categoryServiceImpl) GetCategory(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (c *categoryServiceImpl) UpdateCategory(ctx context.Context, category *model.Category) error {
	panic("unimplemented")
}
