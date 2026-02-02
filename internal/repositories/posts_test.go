package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/gera9/blog/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepositories_CreatePost(t *testing.T) {
	t.Cleanup(func() {
		err := PostgresContainer.Restore(context.TODO())
		require.NoError(t, err)
		PostgresConn.Pool().Reset()
	})

	assertions := assert.New(t)

	type args struct {
		ctx  context.Context
		post models.Post
	}
	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		wantErr bool
		err     error
	}{
		{
			name: "Should return a post",
			args: args{
				ctx: context.TODO(),
				post: models.Post{
					Title:    "Test Post",
					Extract:  "This is a test post",
					Content:  "This is the content of the test post",
					AuthorId: uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertedId, gotErr := PostsRepo.CreatePost(tt.args.ctx, tt.args.post)
			if tt.wantErr {
				assertions.Error(gotErr)
				assertions.Equal(tt.err, gotErr)
				return
			}

			assertions.NoError(gotErr)
			assertions.NotEqual(insertedId, uuid.Nil)
		})
	}
}

func TestRepositories_FindAllPosts(t *testing.T) {
	t.Cleanup(func() {
		err := PostgresContainer.Restore(context.TODO())
		require.NoError(t, err)
		PostgresConn.Pool().Reset()
	})

	assertions := assert.New(t)

	type args struct {
		ctx      context.Context
		limit    int
		offset   int
		authorId uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Post
		wantErr bool
		err     error
	}{
		{
			name: "Should list 10 posts in 1 page (offset 0)",
			args: args{
				ctx:      context.TODO(),
				limit:    10,
				offset:   0,
				authorId: uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
			},
			wantErr: false,
			err:     nil,
			want: []models.Post{
				{
					Id:        uuid.MustParse("91c1538a-518c-4b05-9a1e-180c561a70b3"),
					Title:     "My First Post",
					Extract:   "This is my first post extract.",
					Content:   "This is the full content of my first post.",
					AuthorId:  uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
					CreatedAt: time.Date(2006, 01, 02, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2006, 01, 02, 0, 0, 0, 0, time.UTC),
				},
				{
					Id:        uuid.MustParse("4c09ea12-30ec-4fea-a667-15be9f13e476"),
					Title:     "Another Day in the Life",
					Extract:   "A short story extract.",
					Content:   "A longer text describing my second post.",
					AuthorId:  uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
					CreatedAt: time.Date(2006, 01, 02, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2006, 01, 02, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := PostsRepo.FindAllPosts(tt.args.ctx, tt.args.limit, tt.args.offset, tt.args.authorId)
			if tt.wantErr {
				assertions.Error(gotErr)
				assertions.Equal(tt.err, gotErr)
				return
			}

			assertions.NoError(gotErr)
			assertions.Equal(tt.want, got)
		})
	}
}

func TestRepositories_FindPostByIdAndAuthorId(t *testing.T) {
	t.Cleanup(func() {
		err := PostgresContainer.Restore(context.TODO())
		require.NoError(t, err)
		PostgresConn.Pool().Reset()
	})

	assertions := assert.New(t)

	type args struct {
		ctx      context.Context
		id       uuid.UUID
		authorId uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    models.Post
		wantErr bool
		err     error
	}{
		{
			name: "Should return a post",
			args: args{
				ctx:      context.TODO(),
				id:       uuid.MustParse("91c1538a-518c-4b05-9a1e-180c561a70b3"),
				authorId: uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
			},
			want: models.Post{
				Id:        uuid.MustParse("91c1538a-518c-4b05-9a1e-180c561a70b3"),
				Title:     "My First Post",
				Extract:   "This is my first post extract.",
				Content:   "This is the full content of my first post.",
				AuthorId:  uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
				CreatedAt: time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := PostsRepo.FindPostByIdAndAuthorId(tt.args.ctx, tt.args.id, tt.args.authorId)
			if tt.wantErr {
				if assertions.Error(gotErr) {
					assertions.Equal(tt.err, gotErr, gotErr.Error())
				}
				return
			}

			if !assertions.NoError(gotErr) {
				return
			}
			assertions.Equal(tt.want, got)
		})
	}
}

func TestRepositories_UpdatePostByIdAndAuthorId(t *testing.T) {
	t.Cleanup(func() {
		err := PostgresContainer.Restore(context.TODO())
		require.NoError(t, err)
		PostgresConn.Pool().Reset()
	})

	type args struct {
		ctx      context.Context
		id       uuid.UUID
		authorId uuid.UUID
		post     models.Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update post",
			args: args{
				ctx:      context.TODO(),
				id:       uuid.MustParse("91c1538a-518c-4b05-9a1e-180c561a70b3"),
				authorId: uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
				post: models.Post{
					Id:        uuid.MustParse("91c1538a-518c-4b05-9a1e-180c561a70b3"),
					Title:     "New Title",
					Extract:   "This is my first post extract.",
					Content:   "This is the full content of my first post.",
					AuthorId:  uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
					CreatedAt: time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PostsRepo.UpdatePostByIdAndAuthorId(tt.args.ctx, tt.args.id, tt.args.authorId, tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("Repositories.UpdatePostByIdAndAuthorId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRepositories_DeletePostById(t *testing.T) {
	t.Cleanup(func() {
		err := PostgresContainer.Restore(context.TODO())
		require.NoError(t, err)
		PostgresConn.Pool().Reset()
	})

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
			name: "Should delete post",
			args: args{
				ctx: context.TODO(),
				id:  uuid.MustParse("91c1538a-518c-4b05-9a1e-180c561a70b3"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PostsRepo.DeletePostById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Repositories.DeletePostById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
