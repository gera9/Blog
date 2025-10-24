package services

import (
	"context"

	"github.com/gera9/blog/internal/models"
	"github.com/gera9/blog/pkg/utils"
	"github.com/google/uuid"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, user models.User) (uuid.UUID, error)
	FindAllUsers(ctx context.Context, limit, offset int) ([]models.User, error)
	FindUserById(ctx context.Context, id uuid.UUID) (models.User, error)
	UpdateUserById(ctx context.Context, id uuid.UUID, user models.User) error
	DeleteUserById(ctx context.Context, id uuid.UUID) error
}

func (s Services) CreateUser(ctx context.Context, user models.User) (uuid.UUID, error) {
	return s.UsersRepository.CreateUser(ctx, user)
}

func (s Services) FindAllUsers(ctx context.Context, limit, offset int) ([]models.User, error) {
	return s.UsersRepository.FindAllUsers(ctx, limit, offset)
}

func (s Services) FindUserById(ctx context.Context, id uuid.UUID) (models.User, error) {
	return s.UsersRepository.FindUserById(ctx, id)
}

func (s Services) UpdateUserById(ctx context.Context, id uuid.UUID, newUser models.User) error {
	user, err := s.UsersRepository.FindUserById(ctx, id)
	if err != nil {
		return err
	}

	err = utils.PatchStruct(&user, newUser)
	if err != nil {
		return err
	}

	return s.UsersRepository.UpdateUserById(ctx, id, user)
}

func (s Services) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	return s.UsersRepository.DeleteUserById(ctx, id)
}
