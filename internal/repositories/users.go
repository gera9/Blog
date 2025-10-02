package repositories

import (
	"context"
	"time"

	"github.com/gera9/blog/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserModel struct {
	Id             bson.ObjectID `bson:"_id,omitempty"`
	FirstName      string        `bson:"first_name,omitempty"`
	LastName       string        `bson:"last_name,omitempty"`
	Email          string        `bson:"email,omitempty"`
	Username       string        `bson:"username,omitempty"`
	HashedPassword string        `bson:"hashed_password,omitempty"`
	BirthDate      time.Time     `bson:"birth_date,omitempty"`
	CreatedAt      time.Time     `bson:"created_at,omitempty"`
	UpdatedAt      time.Time     `bson:"updated_at,omitempty"`
}

const usersCollName = "users"

func (r repositories) CreateUser(ctx context.Context, user models.User) (string, error) {
	coll := r.Database().Collection(usersCollName)
	return Create(ctx, coll, toUserModel(user))
}

func (r repositories) FindAllUsers(ctx context.Context, limit, offset int) ([]models.User, error) {
	coll := r.Database().Collection(usersCollName)

	users := []UserModel{}
	err := FindAll(ctx, coll, limit, offset, &users)
	if err != nil {
		return nil, err
	}

	results := make([]models.User, len(users))
	for i, user := range users {
		results[i] = user.toUser()
	}

	return results, nil
}

func (r repositories) FindUserById(ctx context.Context, id string) (models.User, error) {
	coll := r.Database().Collection(usersCollName)

	user := UserModel{}
	err := FindById(ctx, coll, id, &user)
	if err != nil {
		return models.User{}, err
	}

	return user.toUser(), nil
}

func (r repositories) UpdateUserById(ctx context.Context, id string, user models.User) error {
	coll := r.Database().Collection(usersCollName)
	return UpdateById(ctx, coll, id, toUserModel(user))
}

func (r repositories) DeleteUserById(ctx context.Context, id string) error {
	coll := r.Database().Collection(usersCollName)
	return DeleteById(ctx, coll, id)
}

func toUserModel(user models.User) UserModel {
	return UserModel{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		BirthDate:      user.BirthDate,
		UpdatedAt:      user.UpdatedAt,
		CreatedAt:      user.CreatedAt,
	}
}

func (m UserModel) toUser() models.User {
	return models.User{
		Id:             m.Id.Hex(),
		FirstName:      m.FirstName,
		LastName:       m.LastName,
		Email:          m.Email,
		Username:       m.Username,
		HashedPassword: m.HashedPassword,
		BirthDate:      m.BirthDate,
		UpdatedAt:      m.UpdatedAt,
		CreatedAt:      m.CreatedAt,
	}
}
