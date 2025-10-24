package repositories

import (
	"context"
	"time"

	"github.com/gera9/blog/internal/models"
	"github.com/google/uuid"
)

const postsTableName = "posts"

func (r repositories) CreatePost(ctx context.Context, post models.Post) (uuid.UUID, error) {
	now := time.Now().UTC()
	if post.CreatedAt.IsZero() {
		post.CreatedAt = now
	}
	if post.UpdatedAt.IsZero() {
		post.UpdatedAt = now
	}

	sql := `INSERT INTO ` + postsTableName + ` (
		title, extract, content, author_id, created_at, updated_at
	) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	var returnedID uuid.UUID
	err := r.connPool.QueryRow(ctx, sql,
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

func (r repositories) FindAllPosts(ctx context.Context, limit, offset int) ([]models.Post, error) {
	sql := `SELECT id, title, extract, content, author_id, created_at, updated_at
	FROM ` + postsTableName + ` ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := r.connPool.Query(ctx, sql, limit, offset)
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
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r repositories) FindPostById(ctx context.Context, id uuid.UUID) (models.Post, error) {
	sql := `SELECT id, title, extract, content, author_id, created_at, updated_at
	FROM ` + postsTableName + ` WHERE id = $1`

	var post models.Post
	err := r.connPool.QueryRow(ctx, sql, id).Scan(
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

	return post, nil
}

func (r repositories) UpdatePostById(ctx context.Context, id uuid.UUID, post models.Post) error {
	sql := `UPDATE ` + postsTableName + ` SET
		title = $1,
		extract = $2,
		content = $3,
		author_id = $4,
		updated_at = $5
	WHERE id = $6`

	_, err := r.connPool.Exec(ctx, sql,
		post.Title,
		post.Extract,
		post.Content,
		post.AuthorId,
		time.Now().UTC(),
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r repositories) DeletePostById(ctx context.Context, id uuid.UUID) error {
	sql := `DELETE FROM ` + postsTableName + ` WHERE id = $1`

	_, err := r.connPool.Exec(ctx, sql, id)
	if err != nil {
		return err
	}

	return nil
}
