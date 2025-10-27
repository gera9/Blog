package repositories

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	connPool *pgxpool.Pool
}

var (
	once     sync.Once
	instance *Repositories
)

func NewRepositories(ctx context.Context, connStr string) (*Repositories, error) {
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

		instance = &Repositories{connPool}
	})
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (r Repositories) Pool() *pgxpool.Pool {
	return r.connPool
}

func (r Repositories) Close() {
	r.connPool.Close()
}
