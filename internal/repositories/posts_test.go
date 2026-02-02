package repositories_test

import (
	"context"
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

type postsTestsSuite struct {
	suite.Suite
	postsRepo *repositories.PostsRepository
	ctx       context.Context
}

// This will run before running the suite
func (s *postsTestsSuite) SetupSuite() {
	s.postsRepo = repositories.NewPostsRepository(PostgresConn, utils.MockClock{})
}

// This will run after the termination of the suite
func (s *postsTestsSuite) TearDownSuite() {
}

// This will run before each test
func (s *postsTestsSuite) SetupTest() {
}

// This will run after each test
func (s *postsTestsSuite) TearDownTest() {
	err := PostgresContainer.Restore(context.TODO())
	require.NoError(s.T(), err)
	PostgresConn.Pool().Reset()
}

func TestPostsRepoTestSuite(t *testing.T) {
	suite.Run(t, new(postsTestsSuite))
}

func (s *postsTestsSuite) TestCreatePost() {
	t := s.T()

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
			insertedId, gotErr := s.postsRepo.CreatePost(tt.args.ctx, tt.args.post)
			if tt.wantErr {
				assertions.Error(gotErr)
				require.Equal(t, tt.err, gotErr)
			}

			assertions.NoError(gotErr)
			assertions.NotEqual(insertedId, uuid.Nil)
		})
	}
}

func (s *postsTestsSuite) TestFindAllPosts() {
	t := s.T()

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
			got, gotErr := s.postsRepo.FindAllPosts(tt.args.ctx, tt.args.limit, tt.args.offset, tt.args.authorId)
			if tt.wantErr {
				assertions.Error(gotErr)
				require.Equal(t, tt.err, gotErr)
			}

			assertions.NoError(gotErr)
			assertions.Equal(tt.want, got)
		})
	}
}

func (s *postsTestsSuite) TestFindPostByIdAndAuthorId() {
	t := s.T()

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
			got, gotErr := s.postsRepo.FindPostByIdAndAuthorId(tt.args.ctx, tt.args.id, tt.args.authorId)
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

func (s *postsTestsSuite) TestUpdatePostByIdAndAuthorId() {
	t := s.T()

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
			if err := s.postsRepo.UpdatePostByIdAndAuthorId(tt.args.ctx, tt.args.id, tt.args.authorId, tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("Repositories.UpdatePostByIdAndAuthorId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func (s *postsTestsSuite) TestDeletePostById() {
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
			name: "Should delete post",
			args: args{
				ctx: context.TODO(),
				id:  uuid.MustParse("91c1538a-518c-4b05-9a1e-180c561a70b3"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.postsRepo.DeletePostById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Repositories.DeletePostById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
