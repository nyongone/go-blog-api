package usecase

import (
	"context"
	"go-blog-api/domain"
	"go-blog-api/ent"
	"go-blog-api/internal/util"
	"time"
)

type loginUsecase struct {
	userRepository domain.UserRepository
}

func NewLoginUsecase(userRepository domain.UserRepository) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
	}
}

func (u *loginUsecase) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return u.userRepository.GetByEmail(ctx, email)
}

func (u *loginUsecase) CreateAccessToken(user *ent.User) (string, error) {
	return util.CreateAccessToken(user)
}

func (u *loginUsecase) CreateRefreshToken() (string, error) {
	return util.CreateRefreshToken()
}

func (u *loginUsecase) UpdateRefreshToken(ctx context.Context, id int, refreshToken string) error {
	return u.userRepository.UpdateRefreshToken(ctx, id, refreshToken)
}

func (u *loginUsecase) UpdateLastSigninAt(ctx context.Context, id int, date time.Time) error {
	return u.userRepository.UpdateLastSigninAt(ctx, id, date)
}