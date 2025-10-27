package repositories_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/gera9/blog/internal/models"
	"github.com/gera9/blog/internal/repositories"
	"github.com/gera9/blog/internal/repositories/testhelpers"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CustomerRepoTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repository  *repositories.Repositories
	ctx         context.Context
}

func (suite *CustomerRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	pgContainer, err := testhelpers.NewPostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer
	repository, err := repositories.NewRepositories(suite.ctx, suite.pgContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	suite.repository = repository
}

func (suite *CustomerRepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *CustomerRepoTestSuite) TestFindAllUsers() {
	t := suite.T()

	tests := []struct {
		name    string
		limit   int
		offset  int
		wantErr bool
		err     error
		want    []models.User
	}{
		{
			name:    "List 10 users in 1 page (offset 0)",
			limit:   10,
			offset:  0,
			wantErr: false,
			err:     nil,
			want: []models.User{
				{
					Id:             uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
					FirstName:      "Alice",
					LastName:       "Smith",
					Email:          "alice@example.com",
					Username:       "alice_s",
					HashedPassword: "hashed_pwd_1",
					BirthDate:      time.Date(1990, 4, 12, 0, 0, 0, 0, time.UTC),
					CreatedAt:      time.Date(2006, 01, 02, 3, 0, 0, 0, time.Local),
					UpdatedAt:      time.Date(2006, 01, 02, 3, 0, 0, 0, time.Local),
				},
				{
					Id:             uuid.MustParse("b2ccc80d-606e-422f-a9e1-5fd7371163db"),
					FirstName:      "Bob",
					LastName:       "Johnson",
					Email:          "bob@example.com",
					Username:       "bobby_j",
					HashedPassword: "hashed_pwd_2",
					BirthDate:      time.Date(1988, 9, 25, 0, 0, 0, 0, time.UTC),
					CreatedAt:      time.Date(2006, 01, 02, 3, 0, 0, 0, time.Local),
					UpdatedAt:      time.Date(2006, 01, 02, 3, 0, 0, 0, time.Local),
				},
				{
					Id:             uuid.MustParse("2cdc1c8f-9985-4b6c-b007-038a5bef22b5"),
					FirstName:      "Charlie",
					LastName:       "Brown",
					Email:          "charlie@example.com",
					Username:       "charlie_b",
					HashedPassword: "hashed_pwd_3",
					BirthDate:      time.Date(1995, 2, 7, 0, 0, 0, 0, time.UTC),
					CreatedAt:      time.Date(2006, 01, 02, 3, 0, 0, 0, time.Local),
					UpdatedAt:      time.Date(2006, 01, 02, 3, 0, 0, 0, time.Local),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := suite.repository.FindAllUsers(suite.ctx, tt.limit, tt.offset)
			if tt.wantErr {
				if assert.Error(t, gotErr) {
					assert.Equal(t, tt.err, gotErr)
				}
				return
			}

			if assert.NoError(t, gotErr) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func (suite *CustomerRepoTestSuite) TestCreateUser() {
	t := suite.T()

	tests := []struct {
		name    string
		user    models.User
		wantErr bool
		err     error
	}{
		{
			name: "Create user ok",
			user: models.User{
				FirstName:      "Diana",
				LastName:       "Evans",
				Email:          "diana.evans@example.com",
				Username:       "diana_e",
				HashedPassword: "hashed_pwd_4",
				BirthDate:      time.Date(1992, 11, 3, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Create user with duplicated email",
			user: models.User{
				FirstName:      "Diana",
				LastName:       "Evans",
				Email:          "charlie@example.com",
				Username:       "diana_e",
				HashedPassword: "hashed_pwd_4",
				BirthDate:      time.Date(1992, 11, 3, 0, 0, 0, 0, time.UTC),
			},
			wantErr: true,
			err: &pgconn.PgError{
				Severity:            "ERROR",
				SeverityUnlocalized: "ERROR",
				Code:                "23505",
				Message:             "duplicate key value violates unique constraint \"users_email_key\"",
				Detail:              "Key (email)=(charlie@example.com) already exists.",
				SchemaName:          "public",
				TableName:           "users",
				ConstraintName:      "users_email_key",
				File:                "nbtinsert.c",
				Line:                666,
				Routine:             "_bt_check_unique",
			},
		},
		{
			name: "Create user with duplicated username",
			user: models.User{
				FirstName:      "Alice",
				LastName:       "Smith",
				Email:          "alice99@example.com",
				Username:       "alice_s",
				HashedPassword: "hashed_pwd_4",
				BirthDate:      time.Date(1992, 11, 3, 0, 0, 0, 0, time.UTC),
			},
			wantErr: true,
			err: &pgconn.PgError{
				Severity:            "ERROR",
				SeverityUnlocalized: "ERROR",
				Code:                "23505",
				Message:             "duplicate key value violates unique constraint \"users_username_key\"",
				Detail:              "Key (username)=(alice_s) already exists.",
				SchemaName:          "public",
				TableName:           "users",
				ConstraintName:      "users_username_key",
				File:                "nbtinsert.c",
				Line:                666,
				Routine:             "_bt_check_unique",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertedId, gotErr := suite.repository.CreateUser(suite.ctx, tt.user)
			if tt.wantErr {
				if assert.Error(t, gotErr) {
					assert.Equal(t, tt.err, gotErr)
				}
				return
			}

			assert.NotEqual(t, insertedId, uuid.Nil)

			err := suite.repository.DeleteUserById(suite.ctx, insertedId)
			if err != nil {
				t.Fail()
			}
		})
	}
}

func (suite *CustomerRepoTestSuite) TestGetUserrById() {
	t := suite.T()

	customer, err := suite.repository.FindUserById(suite.ctx, uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"))
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, models.User{
		Id:             uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
		FirstName:      "Alice",
		LastName:       "Smith",
		Email:          "alice@example.com",
		Username:       "alice_s",
		HashedPassword: "hashed_pwd_1",
		BirthDate:      time.Date(1990, 4, 12, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Date(2006, 01, 02, 3, 0, 0, 0, time.Local),
		UpdatedAt:      time.Date(2006, 01, 02, 3, 0, 0, 0, time.Local),
	}, customer)
}

func TestCustomerRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepoTestSuite))
}
