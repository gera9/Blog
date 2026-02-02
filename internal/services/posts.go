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

type postsService struct {
	repo PostsRepository
}

func NewPostsService(repo PostsRepository) *postsService {
	return &postsService{repo: repo}
}

func (s postsService) CreatePost(ctx context.Context, post models.Post) (uuid.UUID, error) {
	return s.repo.CreatePost(ctx, post)
}

func (s postsService) FindAllPosts(ctx context.Context, limit, offset int, authorId uuid.UUID) ([]models.Post, error) {
	return s.repo.FindAllPosts(ctx, limit, offset, authorId)
}

func (s postsService) FindPostById(ctx context.Context, id, authorId uuid.UUID) (models.Post, error) {
	return s.repo.FindPostByIdAndAuthorId(ctx, id, authorId)
}

func (s postsService) UpdatePostByIdAndAuthorId(ctx context.Context, id, authorId uuid.UUID, newPost models.Post) error {
	post, err := s.repo.FindPostByIdAndAuthorId(ctx, id, newPost.AuthorId)
	if err != nil {
		return err
	}

	err = utils.PatchStruct(&post, newPost)
	if err != nil {
		return err
	}

	return s.repo.UpdatePostByIdAndAuthorId(ctx, id, authorId, post)
}

func (s postsService) FindPostByIdAndAuthorId(ctx context.Context, id, authorId uuid.UUID) (models.Post, error) {
	return models.Post{}, nil
}

func (s postsService) DeletePostById(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeletePostById(ctx, id)
}
