package domain

import (
	"context"
	"go-blog-api/ent"
)

type PostRepository interface {
	Get(ctx context.Context, id int) (*ent.Post, error)
	GetBySlug(ctx context.Context, slug string) (*ent.Post, error)
	List(ctx context.Context, page int, limit int) ([]*ent.Post, error)
	Create(ctx context.Context, post *ent.Post) error
	Update(ctx context.Context, id int, post *ent.Post) error
	Delete(ctx context.Context, id int) error
}

type PostUsecase interface {
	GetPost(ctx context.Context, id int) (*ent.Post, error)
	GetPostBySlug(ctx context.Context, slug string) (*ent.Post, error)
	GetPostList(ctx context.Context, page int, limit int) ([]*ent.Post, error)
	CreatePost(ctx context.Context, post *ent.Post) error
	UpdatePost(ctx context.Context, id int, post *ent.Post) error
	DeletePost(ctx context.Context, id int) error
}