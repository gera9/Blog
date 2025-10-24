package repositories

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repositories struct {
	connPool *pgxpool.Pool
}

var (
	once     sync.Once
	instance *repositories
)

func NewRepositories(ctx context.Context, connStr string) (*repositories, error) {
	var err error
	once.Do(func() {
		var cfg *pgxpool.Config
		cfg, err = pgxpool.ParseConfig(connStr)
		if err != nil {
			return
		}

		cfg.ConnConfig.TLSConfig = nil

		var connPool *pgxpool.Pool
		connPool, err = pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			return
		}

		err = connPool.Ping(ctx)
		if err != nil {
			return
		}

		instance = &repositories{connPool}
	})
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (r repositories) Close() {
	r.connPool.Close()
}
