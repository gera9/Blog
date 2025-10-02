package repositories

import (
	"context"
	"time"

	"github.com/gera9/blog/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostModel struct {
	Id        bson.ObjectID `bson:"_id,omitempty"`
	Title     string        `bson:"title,omitempty"`
	Extract   string        `bson:"extract,omitempty"`
	Content   string        `bson:"content,omitempty"`
	AuthorId  bson.ObjectID `bson:"author_id,omitempty"`
	CreatedAt time.Time     `bson:"created_at,omitempty"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty"`
}

const postsCollName = "posts"

func (r repositories) CreatePost(ctx context.Context, post models.Post) (string, error) {
	coll := r.Database().Collection(postsCollName)

	m, err := toPostModel(post)
	if err != nil {
		return "", err
	}

	return Create(ctx, coll, m)
}

func (r repositories) FindAllPosts(ctx context.Context, limit, offset int) ([]models.Post, error) {
	coll := r.Database().Collection(postsCollName)

	posts := []PostModel{}
	err := FindAll(ctx, coll, limit, offset, &posts)
	if err != nil {
		return nil, err
	}

	results := make([]models.Post, len(posts))
	for i, post := range posts {
		results[i] = post.toPost()
	}

	return results, nil
}

func (r repositories) FindPostById(ctx context.Context, id string) (models.Post, error) {
	coll := r.Database().Collection(postsCollName)

	post := PostModel{}
	err := FindById(ctx, coll, id, &post)
	if err != nil {
		return models.Post{}, err
	}

	return post.toPost(), nil
}

func (r repositories) UpdatePostById(ctx context.Context, id string, post models.Post) error {
	coll := r.Database().Collection(postsCollName)

	m, err := toPostModel(post)
	if err != nil {
		return err
	}

	return UpdateById(ctx, coll, id, m)
}

func (r repositories) DeletePostById(ctx context.Context, id string) error {
	coll := r.Database().Collection(postsCollName)
	return DeleteById(ctx, coll, id)
}

func toPostModel(post models.Post) (PostModel, error) {
	authorOid, err := bson.ObjectIDFromHex(post.AuthorId)
	if err != nil {
		return PostModel{}, err
	}

	return PostModel{
		Title:     post.Title,
		Extract:   post.Extract,
		Content:   post.Content,
		AuthorId:  authorOid,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

func (m PostModel) toPost() models.Post {
	return models.Post{
		Id:        m.Id.Hex(),
		Title:     m.Title,
		Extract:   m.Extract,
		Content:   m.Content,
		AuthorId:  m.AuthorId.Hex(),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
