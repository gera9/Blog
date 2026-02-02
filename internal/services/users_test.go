package services

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/gera9/blog/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var commonTime = time.Date(2006, time.January, 02, 0, 0, 0, 0, time.UTC)

func TestNewUsersService(t *testing.T) {
	type args struct {
		repo UsersRepository
	}
	tests := []struct {
		name string
		args args
		want *usersService
	}{
		{
			name: "Should create a new user service",
			args: args{
				repo: NewMockUsersRepository(t),
			},
			want: NewUsersService(NewMockUsersRepository(t)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsersService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsersService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_CreateUser(t *testing.T) {
	type fields struct {
		repo *MockUsersRepository
	}
	type args struct {
		ctx  context.Context
		user models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		{
			name: "Should create a new user",
			args: args{
				ctx: context.TODO(),
				user: models.User{
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
			fields: fields{
				repo: func() *MockUsersRepository {
					r := NewMockUsersRepository(t)
					r.On("CreateUser", context.TODO(), models.User{
						Id:             uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
						FirstName:      "Alice",
						LastName:       "Smith",
						Email:          "alice@example.com",
						Username:       "alice_s",
						HashedPassword: "hashed_pwd_1",
						BirthDate:      time.Date(1990, time.April, 12, 0, 0, 0, 0, time.UTC),
						CreatedAt:      commonTime,
						UpdatedAt:      commonTime,
					}).Return(
						uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
						nil,
					)
					return r
				}(),
			},
			want: uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := usersService{
				repo: tt.fields.repo,
			}

			got, err := s.CreateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_usersService_FindAllUsers(t *testing.T) {
	type fields struct {
		repo *MockUsersRepository
	}
	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.User
		wantErr bool
	}{
		{
			name: "Should find all users",
			args: args{
				ctx:    context.TODO(),
				limit:  100,
				offset: 0,
			},
			fields: fields{
				repo: func() *MockUsersRepository {
					r := NewMockUsersRepository(t)
					r.On("FindAllUsers", context.TODO(), 100, 0).Return(
						[]models.User{
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
						},
						nil,
					)
					return r
				}(),
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := usersService{
				repo: tt.fields.repo,
			}

			got, err := s.FindAllUsers(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.FindAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_usersService_FindUserById(t *testing.T) {
	type fields struct {
		repo UsersRepository
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := usersService{
				repo: tt.fields.repo,
			}
			got, err := s.FindUserById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.FindUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersService.FindUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_UpdateUserById(t *testing.T) {
	type fields struct {
		repo UsersRepository
	}
	type args struct {
		ctx     context.Context
		id      uuid.UUID
		newUser models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := usersService{
				repo: tt.fields.repo,
			}
			if err := s.UpdateUserById(tt.args.ctx, tt.args.id, tt.args.newUser); (err != nil) != tt.wantErr {
				t.Errorf("usersService.UpdateUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_usersService_DeleteUserById(t *testing.T) {
	type fields struct {
		repo UsersRepository
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := usersService{
				repo: tt.fields.repo,
			}
			if err := s.DeleteUserById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("usersService.DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
