package repositories_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/gera9/blog/internal/models"
	"github.com/gera9/blog/internal/repositories"
	"github.com/gera9/blog/internal/repositories/testhelpers"
	"github.com/gera9/blog/pkg/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PostsRepoTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repository  *repositories.Repositories
	ctx         context.Context
}

func (suite *PostsRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	pgContainer, err := testhelpers.NewPostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer

	repository, err := repositories.NewRepositories(suite.ctx, suite.pgContainer.ConnectionString, utils.MockClock{})
	if err != nil {
		log.Fatal(err)
	}

	suite.repository = repository
}

func (suite *PostsRepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func TestPostsRepoTestSuite(t *testing.T) {
	suite.Run(t, new(PostsRepoTestSuite))
}

func (suite *PostsRepoTestSuite) TestCreatePost() {
	t := suite.T()

	tests := []struct {
		name    string
		post    models.Post
		wantErr bool
		err     error
	}{
		{
			name: "valid post",
			post: models.Post{
				Title:    "Test Post",
				Extract:  "This is a test post",
				Content:  "This is the content of the test post",
				AuthorId: uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertedId, gotErr := suite.repository.CreatePost(suite.ctx, tt.post)
			if tt.wantErr {
				if assert.Error(t, gotErr) {
					assert.Equal(t, tt.err, gotErr)
				}
				return
			}

			assert.NotEqual(t, insertedId, uuid.Nil)

			err := suite.repository.DeletePostById(suite.ctx, insertedId)
			if err != nil {
				t.Fail()
			}
		})
	}
}

func (suite *PostsRepoTestSuite) TestFindPostById() {
	t := suite.T()

	tests := []struct {
		name     string
		postId   uuid.UUID
		authorId uuid.UUID
		wantErr  bool
		err      error
		want     models.Post
	}{
		{
			name:     "find existing post",
			postId:   uuid.MustParse("91c1538a-518c-4b05-9a1e-180c561a70b3"),
			authorId: uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
			wantErr:  false,
			want: models.Post{
				Id:        uuid.MustParse("91c1538a-518c-4b05-9a1e-180c561a70b3"),
				Title:     "My First Post",
				Extract:   "This is my first post extract.",
				Content:   "This is the full content of my first post.",
				AuthorId:  uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
				CreatedAt: time.Date(2006, 01, 02, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2006, 01, 02, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name:    "find non-existing post",
			postId:  uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			wantErr: true,
			err:     pgx.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := suite.repository.FindPostByIdAndAuthorId(suite.ctx, tt.postId, tt.authorId)
			if tt.wantErr {
				if assert.Error(t, gotErr) {
					assert.Equal(t, tt.err, gotErr)
				}
				return
			}

			assert.NoError(t, gotErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func (suite *PostsRepoTestSuite) TestFindAllPosts() {
	t := suite.T()

	tests := []struct {
		name     string
		limit    int
		offset   int
		authorId uuid.UUID
		wantErr  bool
		err      error
		want     []models.Post
	}{
		{
			name:     "List 10 posts in 1 page (offset 0)",
			limit:    10,
			offset:   0,
			authorId: uuid.MustParse("0853f607-2422-4631-8526-832edaa479c4"),
			wantErr:  false,
			err:      nil,
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
			got, gotErr := suite.repository.FindAllPosts(suite.ctx, tt.limit, tt.offset, tt.authorId)
			if tt.wantErr {
				if assert.Error(t, gotErr) {
					assert.Equal(t, tt.err, gotErr)
				}
				return
			}

			assert.NoError(t, gotErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func (suite *PostsRepoTestSuite) TestDeletePostById() {
	t := suite.T()

	tests := []struct {
		name    string
		postId  uuid.UUID
		wantErr bool
		err     error
	}{
		{
			name:    "delete existing post",
			postId:  uuid.MustParse("d290f1ee-6c54-4b01-90e6-d701748f0851"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := suite.repository.DeletePostById(suite.ctx, tt.postId)
			if tt.wantErr {
				if assert.Error(t, gotErr) {
					assert.Equal(t, tt.err, gotErr)
				}
				return
			}

			assert.NoError(t, gotErr)
		})
	}
}
