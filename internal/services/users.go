package services

import (
	"context"

	"github.com/gera9/blog/internal/models"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, user models.User) (string, error)
	FindAllUsers(ctx context.Context, limit, offset int) ([]models.User, error)
	FindUserById(ctx context.Context, id string) (models.User, error)
	UpdateUserById(ctx context.Context, id string, user models.User) error
	DeleteUserById(ctx context.Context, id string) error
}

func (s Services) CreateUser(ctx context.Context, user models.User) (string, error) {
	return s.UsersRepository.CreateUser(ctx, user)
}

func (s Services) FindAllUsers(ctx context.Context, limit, offset int) ([]models.User, error) {
	return s.UsersRepository.FindAllUsers(ctx, limit, offset)
}

func (s Services) FindUserById(ctx context.Context, id string) (models.User, error) {
	return s.UsersRepository.FindUserById(ctx, id)
}

func (s Services) UpdateUserById(ctx context.Context, id string, user models.User) error {
	return s.UsersRepository.UpdateUserById(ctx, id, user)
}

func (s Services) DeleteUserById(ctx context.Context, id string) error {
	return s.UsersRepository.DeleteUserById(ctx, id)
}
