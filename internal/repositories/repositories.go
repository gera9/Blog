package repositories

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type repositories struct {
	client   *mongo.Client
	database string
}

var (
	once     sync.Once
	instance *repositories
)

func NewRepositories(ctx context.Context, uri, database string) (*repositories, error) {
	var err error
	once.Do(func() {
		clientOpts := options.Client().ApplyURI(uri)

		var client *mongo.Client
		client, err = mongo.Connect(clientOpts)
		if err != nil {
			return
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			return
		}

		instance = &repositories{client, database}
	})
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (r repositories) Database() *mongo.Database {
	return r.client.Database(r.database)
}
