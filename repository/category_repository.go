package repository

import (
	"context"
	"go-blog-api/domain"
	"go-blog-api/ent"
	"go-blog-api/ent/category"
)

type categoryRepository struct {
	client *ent.Client
}

func NewCategoryRepository(client *ent.Client) domain.CategoryRepository {
	return &categoryRepository{
		client: client,
	}
}

func (r *categoryRepository) Get(ctx context.Context, id int) (*ent.Category, error) {
	category, err := r.client.Category.
									Query().
									Where(category.IDEQ(id)).
									Only(ctx)
	
	if err != nil {
		return nil, err
	}

	return category, nil
} 

func (r *categoryRepository) List(ctx context.Context) ([]*ent.Category, error) {
	categories, err := r.client.Category.
										Query().
										All(ctx)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) Create(ctx context.Context, category *ent.Category) error {
	_, err := r.client.Category.
						Create().
						SetName(category.Name).
						SetSlug(category.Slug).
						Save(ctx)
	
	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) Update(ctx context.Context, id int, category *ent.Category) error {
	_, err := r.client.Category.
						UpdateOneID(id).
						SetName(category.Name).
						SetSlug(category.Slug).
						Save(ctx)
	
	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) Delete(ctx context.Context, id int) error {
	err := r.client.Category.
				DeleteOneID(id).
				Exec(ctx)
				
	if err != nil {
		return err
	}

	return nil
}