package models

import "time"

type User struct {
	Id             string
	FirstName      string
	LastName       string
	Email          string
	Username       string
	HashedPassword string
	BirthDate      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
