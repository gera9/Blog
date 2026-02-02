package postgres

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	connPool *pgxpool.Pool
}

var (
	once     sync.Once
	instance *Postgres
)

func NewPostgres(ctx context.Context, connStr string) (*Postgres, error) {
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

		instance = &Postgres{connPool}
	})
	if err != nil {
		return nil, err
	}

	return instance, nil
}

// Get underlying connections pool.
func (p Postgres) Pool() *pgxpool.Pool {
	return p.connPool
}

func (p Postgres) Close() {
	p.connPool.Close()
}
