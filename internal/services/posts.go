package services

import (
	"context"

	"github.com/gera9/blog/internal/models"
)

type PostsRepository interface {
	CreatePost(ctx context.Context, post models.Post) (string, error)
	FindAllPosts(ctx context.Context, limit, offset int) ([]models.Post, error)
	FindPostById(ctx context.Context, id string) (models.Post, error)
	UpdatePostById(ctx context.Context, id string, post models.Post) error
	DeletePostById(ctx context.Context, id string) error
}

func (s Services) CreatePost(ctx context.Context, post models.Post) (string, error) {
	return s.PostsRepository.CreatePost(ctx, post)
}

func (s Services) FindAllPosts(ctx context.Context, limit, offset int) ([]models.Post, error) {
	return s.PostsRepository.FindAllPosts(ctx, limit, offset)
}

func (s Services) FindPostById(ctx context.Context, id string) (models.Post, error) {
	return s.PostsRepository.FindPostById(ctx, id)
}

func (s Services) UpdatePostById(ctx context.Context, id string, post models.Post) error {
	return s.PostsRepository.UpdatePostById(ctx, id, post)
}

func (s Services) DeletePostById(ctx context.Context, id string) error {
	return s.PostsRepository.DeletePostById(ctx, id)
}
