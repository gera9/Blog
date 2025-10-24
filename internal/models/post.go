package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id        uuid.UUID
	Title     string
	Extract   string
	Content   string
	AuthorId  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
