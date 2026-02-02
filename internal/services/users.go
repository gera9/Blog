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

type usersService struct {
	repo UsersRepository
}

func NewUsersService(repo UsersRepository) *usersService {
	return &usersService{repo: repo}
}

func (s usersService) CreateUser(ctx context.Context, user models.User) (uuid.UUID, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s usersService) FindAllUsers(ctx context.Context, limit, offset int) ([]models.User, error) {
	return s.repo.FindAllUsers(ctx, limit, offset)
}

func (s usersService) FindUserById(ctx context.Context, id uuid.UUID) (models.User, error) {
	return s.repo.FindUserById(ctx, id)
}

func (s usersService) UpdateUserById(ctx context.Context, id uuid.UUID, newUser models.User) error {
	user, err := s.repo.FindUserById(ctx, id)
	if err != nil {
		return err
	}

	err = utils.PatchStruct(&user, newUser)
	if err != nil {
		return err
	}

	return s.repo.UpdateUserById(ctx, id, user)
}

func (s usersService) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteUserById(ctx, id)
}
