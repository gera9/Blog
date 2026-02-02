package repositories

import (
	"context"

	"github.com/gera9/blog/internal/models"
	"github.com/gera9/blog/pkg/postgres"
	"github.com/gera9/blog/pkg/utils"
	"github.com/google/uuid"
)

type PostsRepository struct {
	conn         *postgres.Postgres
	timeProvider utils.TimeProvider
	tableName    string
}

func NewPostsRepository(conn *postgres.Postgres, timeProvider utils.TimeProvider) *PostsRepository {
	return &PostsRepository{
		conn:         conn,
		timeProvider: timeProvider,
		tableName:    "posts",
	}
}

func (r PostsRepository) CreatePost(ctx context.Context, post models.Post) (uuid.UUID, error) {
	now := r.timeProvider.Now().UTC()
	if post.CreatedAt.IsZero() {
		post.CreatedAt = now
	}
	if post.UpdatedAt.IsZero() {
		post.UpdatedAt = now
	}

	sql := `INSERT INTO ` + r.tableName + ` (
		title, extract, content, author_id, created_at, updated_at
	) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	var returnedID uuid.UUID
	err := r.conn.Pool().QueryRow(ctx, sql,
		post.Title,
		post.Extract,
		post.Content,
		post.AuthorId,
		post.CreatedAt,
		post.UpdatedAt,
	).Scan(&returnedID)
	if err != nil {
		return uuid.Nil, err
	}

	return returnedID, nil
}

func (r PostsRepository) FindAllPosts(ctx context.Context, limit, offset int, authorId uuid.UUID) ([]models.Post, error) {
	sql := `SELECT id, title, extract, content, author_id, created_at, updated_at
	FROM ` + r.tableName + ` WHERE author_id = $3 ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := r.conn.Pool().Query(ctx, sql, limit, offset, authorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]models.Post, 0)
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Extract,
			&post.Content,
			&post.AuthorId,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		post.CreatedAt = post.CreatedAt.UTC()
		post.UpdatedAt = post.UpdatedAt.UTC()

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r PostsRepository) FindPostByIdAndAuthorId(ctx context.Context, id, authorId uuid.UUID) (models.Post, error) {
	sql := `SELECT id, title, extract, content, author_id, created_at, updated_at
	FROM ` + r.tableName + ` WHERE id = $1 AND author_id = $2`

	var post models.Post
	err := r.conn.Pool().QueryRow(ctx, sql, id, authorId).Scan(
		&post.Id,
		&post.Title,
		&post.Extract,
		&post.Content,
		&post.AuthorId,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return models.Post{}, err
	}

	post.CreatedAt = post.CreatedAt.UTC()
	post.UpdatedAt = post.UpdatedAt.UTC()

	return post, nil
}

func (r PostsRepository) UpdatePostByIdAndAuthorId(ctx context.Context, id, authorId uuid.UUID, post models.Post) error {
	sql := `UPDATE ` + r.tableName + ` SET
		title = $1,
		extract = $2,
		content = $3,
		author_id = $4,
		updated_at = $5
	WHERE id = $6 AND author_id = $7`

	_, err := r.conn.Pool().Exec(ctx, sql,
		post.Title,
		post.Extract,
		post.Content,
		post.AuthorId,
		r.timeProvider.Now().UTC(),
		id,
		authorId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r PostsRepository) DeletePostById(ctx context.Context, id uuid.UUID) error {
	sql := `DELETE FROM ` + r.tableName + ` WHERE id = $1`

	_, err := r.conn.Pool().Exec(ctx, sql, id)
	if err != nil {
		return err
	}

	return nil
}
