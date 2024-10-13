package db

import (
	"context"

	"github.com/KOBATATU/todo/ent"
)


type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{client: client}
}

// ユーザ作成
func (r *UserRepository) CreateUser(ctx context.Context, email, password string) (*ent.User, error) {
	return r.client.User.Create().
		SetEmail(email).
		SetPassword(password).
		Save(ctx)
}