package domain

import (
	"context"
	"go-blog-api/ent"
)

type UserRepository interface {
	Get(ctx context.Context, id int) (*ent.User, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	Create(ctx context.Context, user *ent.User) error
	Update(ctx context.Context, id int, user *ent.User) error
	UpdateRefreshToken(ctx context.Context, id int, refreshToken string) error
	Delete(ctx context.Context, id int) error
}