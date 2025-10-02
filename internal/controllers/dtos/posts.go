package dtos

import (
	"time"

	"github.com/gera9/blog/internal/models"
)

type CreatePost struct {
	Title    string `json:"title"`
	Extract  string `json:"extract"`
	Content  string `json:"content"`
	AuthorId string `json:"author_id"`
}

func (cp CreatePost) ToPost() models.Post {
	return models.Post{
		Title:    cp.Title,
		Extract:  cp.Extract,
		Content:  cp.Content,
		AuthorId: cp.AuthorId,
	}
}

type UpdatePost struct {
	Title   string `json:"title"`
	Extract string `json:"extract"`
	Content string `json:"content"`
}

func (up UpdatePost) ToPost() models.Post {
	return models.Post{
		Title:   up.Title,
		Extract: up.Extract,
		Content: up.Content,
	}
}

type PostResponse struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Extract   string    `json:"extract"`
	Content   string    `json:"content"`
	AuthorId  string    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
