package dtos

import (
	"time"

	"github.com/gera9/blog/internal/models"
)

type CreateUser struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
}

func (cu CreateUser) ToUser() models.User {
	return models.User{
		FirstName:      cu.FirstName,
		LastName:       cu.LastName,
		Email:          cu.Email,
		Username:       cu.Username,
		HashedPassword: cu.Password,
		BirthDate:      cu.BirthDate,
	}
}

type UpdateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (uu UpdateUser) ToUser() models.User {
	return models.User{
		FirstName:      uu.FirstName,
		LastName:       uu.LastName,
		Email:          uu.Email,
		Username:       uu.Username,
		HashedPassword: uu.Password,
	}
}

type UserResponse struct {
	Id             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	BirthDate      time.Time `json:"birth_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
