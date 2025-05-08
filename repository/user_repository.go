package repository

import (
	"context"
	"go-blog-api/domain"
	"go-blog-api/ent"
	"go-blog-api/ent/user"
	"strings"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) domain.UserRepository {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) Get(ctx context.Context, id int) (*ent.User, error) {
	user, err := r.client.User.
							Query().
							Where(user.IDEQ(id)).
							Only(ctx)
	
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	user, err := r.client.User.
							Query().
							Where(user.EmailEQ(email)).
							Only(ctx)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *userRepository) Create(ctx context.Context, user *ent.User) error {
	_, err := r.client.User.
						Create().
						SetEmail(user.Email).
						SetPassword(user.Password).
						SetName(user.Name).
						Save(ctx)

	if err != nil {
		return err 
	}

	return nil
}

func (r *userRepository) Update(ctx context.Context, id int, user *ent.User) error {
	query := r.client.User.
					UpdateOneID(id)

	if strings.TrimSpace(user.Password) != "" {
		query = query.SetPassword(user.Password)
	}

	_, err := query.
						SetEmail(user.Email).
						SetName(user.Name).
						Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateRefreshToken(ctx context.Context, id int, refreshToken string) error {
	_, err := r.client.User.
						UpdateOneID(id).
						SetRefreshToken(refreshToken).
						Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	err := r.client.User.
				DeleteOneID(id).
				Exec(ctx)
	
	if err != nil {
		return err
	}

	return nil
}