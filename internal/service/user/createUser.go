package service

import (
	"context"

	"github.com/KOBATATU/todo/ent"
	db "github.com/KOBATATU/todo/internal/repository/db/user"
)



type UserService struct {
	repo *db.UserRepository
}

func NewUserService(repo *db.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, email, password string) (*ent.User, error) {
	return s.repo.CreateUser(ctx, email, password)
}
