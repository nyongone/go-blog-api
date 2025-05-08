package domain

import (
	"context"
	"go-blog-api/ent"
)

type CategoryRepository interface {
	Get(ctx context.Context, id int) (*ent.Category, error)
	List(ctx context.Context) ([]*ent.Category, error)
	Create(ctx context.Context, category *ent.Category) error
	Update(ctx context.Context, id int, category *ent.Category) error
	Delete(ctx context.Context, id int) error
}

type CategoryUsecase interface {
	GetCategory(ctx context.Context, id int) (*ent.Category, error)
	GetCategoryList(ctx context.Context) ([]*ent.Category, error)
	CreateCategory(ctx context.Context, category *ent.Category) error
	UpdateCategory(ctx context.Context, id int, category *ent.Category) error
	DeleteCategory(ctx context.Context, id int) error
}