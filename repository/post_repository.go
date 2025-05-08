package repository

import (
	"context"
	"go-blog-api/domain"
	"go-blog-api/ent"
	"go-blog-api/ent/post"

	"entgo.io/ent/dialect/sql"
)

type postRepository struct {
	client	*ent.Client
}

func NewPostRepository(client *ent.Client) domain.PostRepository {
	return &postRepository{
		client: client,
	}
}

func (r *postRepository) Get(ctx context.Context, id int) (*ent.Post, error) {
	post, err := r.client.Post.
							Query().
							Where(post.IDEQ(id)).
							Only(ctx)		

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) GetBySlug(ctx context.Context, slug string) (*ent.Post, error) {
	post, err := r.client.Post.
							Query().
							Where(post.SlugEQ(slug)).
							Only(ctx)
		
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) List(ctx context.Context, page int, limit int) ([]*ent.Post, error) {
	posts, err := r.client.Post.
								Query().
								Limit(limit).
								Offset((page - 1) * limit).
								Order(
									post.ByCreatedAt(
										sql.OrderAsc(),
									),
								).
								All(ctx)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *postRepository) Create(ctx context.Context, post *ent.Post) error {
	_, err := r.client.Post.
						Create().
						SetTitle(post.Title).
						SetSlug(post.Slug).
						SetContent(post.Content).
						SetCategoryID(post.CategoryId).
						Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *postRepository) Update(ctx context.Context, id int, post *ent.Post) error {
	_, err := r.client.Post.
						UpdateOneID(id).
						SetTitle(post.Title).
						SetSlug(post.Slug).
						SetContent(post.Content).
						SetCategoryID(post.CategoryId).
						Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *postRepository) Delete(ctx context.Context, id int) error {
	err := r.client.Post.
						DeleteOneID(id).
						Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}