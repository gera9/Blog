package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/gera9/blog/internal/models"
	"github.com/google/uuid"
)

func TestNewUsersService(t *testing.T) {
	type args struct {
		repo UsersRepository
	}
	tests := []struct {
		name string
		args args
		want *usersService
	}{
		// TODO: Add test cases.
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
		repo UsersRepository
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
		// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_FindAllUsers(t *testing.T) {
	type fields struct {
		repo UsersRepository
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
		// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersService.FindAllUsers() = %v, want %v", got, tt.want)
			}
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
