package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/gera9/blog/internal/models"
	"github.com/gera9/blog/pkg/postgres"
	"github.com/gera9/blog/pkg/utils"
	"github.com/google/uuid"
)

const usersTableName = "users"

type UsersRepository struct {
	conn         *postgres.Postgres
	timeProvider utils.TimeProvider
	tableName    string
}

func NewUsersRepository(conn *postgres.Postgres, timeProvider utils.TimeProvider) *UsersRepository {
	return &UsersRepository{
		conn:         conn,
		timeProvider: timeProvider,
		tableName:    "users",
	}
}

func (r UsersRepository) CreateUser(ctx context.Context, user models.User) (uuid.UUID, error) {
	now := r.timeProvider.Now().UTC()
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = now
	}

	sql := `INSERT INTO ` + usersTableName + ` (
		first_name, last_name, email, username, hashed_password, birth_date, created_at, updated_at
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`

	var returnedID uuid.UUID
	err := r.conn.Pool().QueryRow(ctx, sql,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Username,
		user.HashedPassword,
		user.BirthDate,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&returnedID)
	if err != nil {
		return uuid.Nil, err
	}

	return returnedID, nil
}

func (r UsersRepository) FindAllUsers(ctx context.Context, limit, offset int) ([]models.User, error) {
	sql := `SELECT id, first_name, last_name, email, username, hashed_password, birth_date, created_at, updated_at
	FROM ` + usersTableName + ` ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := r.conn.Pool().Query(ctx, sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var (
			id         uuid.UUID
			firstName  string
			lastName   string
			email      string
			username   string
			hashedPass string
			birthDate  time.Time
			createdAt  time.Time
			updatedAt  time.Time
		)

		if err := rows.Scan(&id, &firstName, &lastName, &email, &username, &hashedPass, &birthDate, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		createdAt = createdAt.UTC()
		updatedAt = updatedAt.UTC()

		users = append(users, models.User{
			Id:             id,
			FirstName:      firstName,
			LastName:       lastName,
			Email:          email,
			Username:       username,
			HashedPassword: hashedPass,
			BirthDate:      birthDate,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
		})
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}

func (r UsersRepository) FindUserById(ctx context.Context, id uuid.UUID) (models.User, error) {
	sql := `SELECT id, first_name, last_name, email, username, hashed_password, birth_date, created_at, updated_at
	FROM ` + usersTableName + ` WHERE id = $1`

	var (
		uuid       uuid.UUID
		firstName  string
		lastName   string
		email      string
		username   string
		hashedPass string
		birthDate  time.Time
		createdAt  time.Time
		updatedAt  time.Time
	)

	err := r.conn.Pool().QueryRow(ctx, sql, id).Scan(&uuid, &firstName, &lastName, &email, &username, &hashedPass, &birthDate, &createdAt, &updatedAt)
	if err != nil {
		return models.User{}, err
	}

	createdAt = createdAt.UTC()
	updatedAt = updatedAt.UTC()

	return models.User{
		Id:             uuid,
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		Username:       username,
		HashedPassword: hashedPass,
		BirthDate:      birthDate,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}, nil
}

func (r UsersRepository) UpdateUserById(ctx context.Context, id uuid.UUID, user models.User) error {
	// update the allowed fields and updated_at
	now := r.timeProvider.Now().UTC()

	sql := `UPDATE ` + usersTableName + ` SET
		first_name = $1,
		last_name = $2,
		email = $3,
		username = $4,
		hashed_password = $5,
		updated_at = $6
	WHERE id = $7`

	tag, err := r.conn.Pool().Exec(ctx, sql,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Username,
		user.HashedPassword,
		now,
		id,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("no rows updated")
	}

	return nil
}

func (r UsersRepository) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	sql := `DELETE FROM ` + usersTableName + ` WHERE id = $1`

	tag, err := r.conn.Pool().Exec(ctx, sql, id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("no rows deleted")
	}

	return nil
}
