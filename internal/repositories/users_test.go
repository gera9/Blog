package repositories_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/gera9/blog/internal/models"
	"github.com/gera9/blog/internal/repositories"
	"github.com/gera9/blog/pkg/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type usersTestsSuite struct {
	suite.Suite
	usersRepo *repositories.UsersRepository
	ctx       context.Context
}

// This will run before running the suite
func (s *usersTestsSuite) SetupSuite() {
	s.usersRepo = repositories.NewUsersRepository(PostgresConn, utils.MockClock{})
}

// This will run after the termination of the suite
func (s *usersTestsSuite) TearDownSuite() {
}

// This will run before each test
func (s *usersTestsSuite) SetupTest() {
}

// This will run after each test
func (s *usersTestsSuite) TearDownTest() {
	err := PostgresContainer.Restore(context.TODO())
	require.NoError(s.T(), err)
	PostgresConn.Pool().Reset()
}

func TestUsersRepoTestSuite(t *testing.T) {
	suite.Run(t, new(usersTestsSuite))
}

var commonTime = time.Date(2006, time.January, 02, 0, 0, 0, 0, time.UTC)

func (s *usersTestsSuite) TestCreateUser() {
	t := s.T()

	assertions := assert.New(t)

	type args struct {
		ctx  context.Context
		user models.User
	}
	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		wantErr bool
		err     error
	}{
		{
			name: "Should create an user",
			args: args{
				ctx: context.TODO(),
				user: models.User{
					FirstName:      "Jane",
					LastName:       "Doe",
					Email:          "jane.doe@example.com",
					Username:       "jdoe_90",
					HashedPassword: "$2a$12$mqL0eS.17/6.8q7/8.0..0.8.q8/7.0.8.q8/7.0.8.q8/7.0",
					BirthDate:      time.Date(1990, time.May, 15, 0, 0, 0, 0, time.UTC),
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertedId, gotErr := s.usersRepo.CreateUser(tt.args.ctx, tt.args.user)
			if tt.wantErr {
				assertions.Error(gotErr)
				require.Equal(t, tt.err, gotErr)
			}

			assertions.NoError(gotErr)
			assertions.NotEqual(insertedId, uuid.Nil)
		})
	}
}

func (s *usersTestsSuite) TestFindAllUsers() {
	t := s.T()

	assertions := assert.New(t)

	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    []models.User
		wantErr bool
		err     error
	}{
		{
			name: "Should return a list of users",
			args: args{
				ctx:    context.TODO(),
				limit:  100,
				offset: 0,
			},
			want: []models.User{
				{
					Id:             uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
					FirstName:      "Alice",
					LastName:       "Smith",
					Email:          "alice@example.com",
					Username:       "alice_s",
					HashedPassword: "hashed_pwd_1",
					BirthDate:      time.Date(1990, time.April, 12, 0, 0, 0, 0, time.UTC),
					CreatedAt:      commonTime,
					UpdatedAt:      commonTime,
				},
				{
					Id:             uuid.MustParse("b2ccc80d-606e-422f-a9e1-5fd7371163db"),
					FirstName:      "Bob",
					LastName:       "Johnson",
					Email:          "bob@example.com",
					Username:       "bobby_j",
					HashedPassword: "hashed_pwd_2",
					BirthDate:      time.Date(1988, time.September, 25, 0, 0, 0, 0, time.UTC),
					CreatedAt:      commonTime,
					UpdatedAt:      commonTime,
				},
				{
					Id:             uuid.MustParse("2cdc1c8f-9985-4b6c-b007-038a5bef22b5"),
					FirstName:      "Charlie",
					LastName:       "Brown",
					Email:          "charlie@example.com",
					Username:       "charlie_b",
					HashedPassword: "hashed_pwd_3",
					BirthDate:      time.Date(1995, time.February, 7, 0, 0, 0, 0, time.UTC),
					CreatedAt:      commonTime,
					UpdatedAt:      commonTime,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := s.usersRepo.FindAllUsers(tt.args.ctx, tt.args.limit, tt.args.offset)
			if tt.wantErr {
				assertions.Error(gotErr)
				require.Equal(t, tt.err, gotErr)
				return
			}

			assertions.NoError(gotErr)
			assertions.Equal(tt.want, got)
		})
	}
}

func (s *usersTestsSuite) TestFindUserById() {
	t := s.T()

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    models.User
		wantErr bool
	}{
		{
			name: "Should find an user",
			args: args{
				ctx: context.TODO(),
				id:  uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
			},
			want: models.User{
				Id:             uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
				FirstName:      "Alice",
				LastName:       "Smith",
				Email:          "alice@example.com",
				Username:       "alice_s",
				HashedPassword: "hashed_pwd_1",
				BirthDate:      time.Date(1990, time.April, 12, 0, 0, 0, 0, time.UTC),
				CreatedAt:      commonTime,
				UpdatedAt:      commonTime,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.usersRepo.FindUserById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repositories.FindUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repositories.FindUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *usersTestsSuite) TestUpdateUserById() {
	t := s.T()

	type args struct {
		ctx  context.Context
		id   uuid.UUID
		user models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update user",
			args: args{
				ctx: context.TODO(),
				id:  uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
				user: models.User{
					Id:             uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
					FirstName:      "Alice",
					LastName:       "Smith",
					Email:          "alice@gmail.com",
					Username:       "alice_s",
					HashedPassword: "hashed_pwd_1",
					BirthDate:      time.Date(1990, time.April, 12, 0, 0, 0, 0, time.UTC),
					CreatedAt:      commonTime,
					UpdatedAt:      commonTime,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.usersRepo.UpdateUserById(tt.args.ctx, tt.args.id, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Repositories.UpdateUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *usersTestsSuite) TestDeleteUserById() {
	t := s.T()

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should delete user",
			args: args{
				ctx: context.TODO(),
				id:  uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.usersRepo.DeleteUserById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Repositories.DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
