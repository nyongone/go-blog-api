package domain

import (
	"context"
	"go-blog-api/ent"
	"time"
)

type LoginRequest struct {
	Email			string		`json:"email"`
	Password	string		`json:"password"`
}

type LoginResponse struct {
	AccessToken		string	`json:"accessToken"`
	RefreshToken	string	`json:"refreshToken"`
}
 
type LoginUsecase interface {
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	CreateAccessToken(user *ent.User) (string, error)
	CreateRefreshToken() (string, error)
	UpdateRefreshToken(ctx context.Context, id int, refreshToken string) error
	UpdateLastSigninAt(ctx context.Context, id int, date time.Time) error
}