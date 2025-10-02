package models

import "time"

type Post struct {
	Id        string
	Title     string
	Extract   string
	Content   string
	AuthorId  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
