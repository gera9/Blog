package services

import (
	"context"

	"github.com/gera9/blog/internal/models"
	"github.com/gera9/blog/pkg/utils"
	"github.com/google/uuid"
)

type PostsRepository interface {
	CreatePost(ctx context.Context, post models.Post) (uuid.UUID, error)
	FindAllPosts(ctx context.Context, limit, offset int, authorId uuid.UUID) ([]models.Post, error)
	FindPostByIdAndAuthorId(ctx context.Context, id, authorId uuid.UUID) (models.Post, error)
	UpdatePostByIdAndAuthorId(ctx context.Context, id, authorId uuid.UUID, post models.Post) error
	DeletePostById(ctx context.Context, id uuid.UUID) error
}

func (s Services) CreatePost(ctx context.Context, post models.Post) (uuid.UUID, error) {
	return s.PostsRepository.CreatePost(ctx, post)
}

func (s Services) FindAllPosts(ctx context.Context, limit, offset int, authorId uuid.UUID) ([]models.Post, error) {
	return s.PostsRepository.FindAllPosts(ctx, limit, offset, authorId)
}

func (s Services) FindPostById(ctx context.Context, id, authorId uuid.UUID) (models.Post, error) {
	return s.PostsRepository.FindPostByIdAndAuthorId(ctx, id, authorId)
}

func (s Services) UpdatePostByIdAndAuthorId(ctx context.Context, id, authorId uuid.UUID, newPost models.Post) error {
	post, err := s.PostsRepository.FindPostByIdAndAuthorId(ctx, id, newPost.AuthorId)
	if err != nil {
		return err
	}

	err = utils.PatchStruct(&post, newPost)
	if err != nil {
		return err
	}

	return s.PostsRepository.UpdatePostByIdAndAuthorId(ctx, id, authorId, post)
}

func (s Services) DeletePostById(ctx context.Context, id uuid.UUID) error {
	return s.PostsRepository.DeletePostById(ctx, id)
}
