package services

import (
	"context"

	"github.com/gera9/blog/internal/models"
	"github.com/gera9/blog/pkg/utils"
	"github.com/google/uuid"
)

type PostsRepository interface {
	CreatePost(ctx context.Context, post models.Post) (uuid.UUID, error)
	FindAllPosts(ctx context.Context, limit, offset int) ([]models.Post, error)
	FindPostById(ctx context.Context, id uuid.UUID) (models.Post, error)
	UpdatePostById(ctx context.Context, id uuid.UUID, post models.Post) error
	DeletePostById(ctx context.Context, id uuid.UUID) error
}

func (s Services) CreatePost(ctx context.Context, post models.Post) (uuid.UUID, error) {
	return s.PostsRepository.CreatePost(ctx, post)
}

func (s Services) FindAllPosts(ctx context.Context, limit, offset int) ([]models.Post, error) {
	return s.PostsRepository.FindAllPosts(ctx, limit, offset)
}

func (s Services) FindPostById(ctx context.Context, id uuid.UUID) (models.Post, error) {
	return s.PostsRepository.FindPostById(ctx, id)
}

func (s Services) UpdatePostById(ctx context.Context, id uuid.UUID, newPost models.Post) error {
	post, err := s.PostsRepository.FindPostById(ctx, id)
	if err != nil {
		return err
	}

	err = utils.PatchStruct(&post, newPost)
	if err != nil {
		return err
	}

	return s.PostsRepository.UpdatePostById(ctx, id, post)
}

func (s Services) DeletePostById(ctx context.Context, id uuid.UUID) error {
	return s.PostsRepository.DeletePostById(ctx, id)
}
