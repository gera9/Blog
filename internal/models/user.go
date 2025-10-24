package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID
	FirstName      string
	LastName       string
	Email          string
	Username       string
	HashedPassword string
	BirthDate      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
