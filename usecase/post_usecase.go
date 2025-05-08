package usecase

import (
	"context"
	"go-blog-api/domain"
	"go-blog-api/ent"
)

type postUsecase struct {
	postRepository domain.PostRepository
}

func NewPostUsecase(postRepository domain.PostRepository) domain.PostUsecase {
	return &postUsecase{
		postRepository: postRepository,
	}
}

func (u *postUsecase) GetPost(ctx context.Context, id int) (*ent.Post, error) {
	return u.postRepository.Get(ctx, id)
}

func (u *postUsecase) GetPostBySlug(ctx context.Context, slug string) (*ent.Post, error) {
	return u.postRepository.GetBySlug(ctx, slug)
}

func (u *postUsecase) GetPostList(ctx context.Context, page int, limit int) ([]*ent.Post, error) {
	return u.postRepository.List(ctx, page, limit)
}

func (u *postUsecase) CreatePost(ctx context.Context, post *ent.Post) error {
	return u.postRepository.Create(ctx, post)
}

func (u *postUsecase) UpdatePost(ctx context.Context, id int, post *ent.Post) error {
	return u.postRepository.Update(ctx, id, post)
}

func (u *postUsecase) DeletePost(ctx context.Context, id int) error {
	return u.postRepository.Delete(ctx, id)
}