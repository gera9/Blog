package repositories

import (
	"context"

	"github.com/gera9/blog/internal/models"
	"github.com/google/uuid"
)

const postsTableName = "posts"

func (r Repositories) CreatePost(ctx context.Context, post models.Post) (uuid.UUID, error) {
	now := r.timeProvider.Now().UTC()
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

func (r Repositories) FindAllPosts(ctx context.Context, limit, offset int, authorId uuid.UUID) ([]models.Post, error) {
	sql := `SELECT id, title, extract, content, author_id, created_at, updated_at
	FROM ` + postsTableName + ` WHERE author_id = $3 ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := r.connPool.Query(ctx, sql, limit, offset, authorId)
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

func (r Repositories) FindPostByIdAndAuthorId(ctx context.Context, id, authorId uuid.UUID) (models.Post, error) {
	sql := `SELECT id, title, extract, content, author_id, created_at, updated_at
	FROM ` + postsTableName + ` WHERE id = $1 AND author_id = $2`

	var post models.Post
	err := r.connPool.QueryRow(ctx, sql, id, authorId).Scan(
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

func (r Repositories) UpdatePostByIdAndAuthorId(ctx context.Context, id, authorId uuid.UUID, post models.Post) error {
	sql := `UPDATE ` + postsTableName + ` SET
		title = $1,
		extract = $2,
		content = $3,
		author_id = $4,
		updated_at = $5
	WHERE id = $6 AND author_id = $7`

	_, err := r.connPool.Exec(ctx, sql,
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

func (r Repositories) DeletePostById(ctx context.Context, id uuid.UUID) error {
	sql := `DELETE FROM ` + postsTableName + ` WHERE id = $1`

	_, err := r.connPool.Exec(ctx, sql, id)
	if err != nil {
		return err
	}

	return nil
}
