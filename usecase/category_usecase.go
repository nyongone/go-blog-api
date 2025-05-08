package usecase

import (
	"context"
	"go-blog-api/domain"
	"go-blog-api/ent"
)

type categoryUsecase struct {
	categoryRepository	domain.CategoryRepository
}

func NewCategoryUsecase(categoryRepository domain.CategoryRepository) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepository: categoryRepository,
	}
}

func (u *categoryUsecase) GetCategory(ctx context.Context, id int) (*ent.Category, error) {
	return u.categoryRepository.Get(ctx, id)
}

func (u *categoryUsecase) GetCategoryList(ctx context.Context) ([]*ent.Category, error) {
	return u.categoryRepository.List(ctx)
}

func (u *categoryUsecase) CreateCategory(ctx context.Context, category *ent.Category) error {
	return u.categoryRepository.Create(ctx, category)
}

func (u *categoryUsecase) UpdateCategory(ctx context.Context, id int, category *ent.Category) error {
	return u.categoryRepository.Update(ctx, id, category)
}

func (u *categoryUsecase) DeleteCategory(ctx context.Context, id int) error {
	return u.categoryRepository.Delete(ctx, id)
}