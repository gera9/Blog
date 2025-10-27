package testhelpers

import (
	"context"
	"path/filepath"

	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

func NewPostgresContainer(ctx context.Context) (*PostgresContainer, error) {
	pgContainer, err := postgres.Run(ctx, "postgres:17-alpine",
		postgres.WithDatabase("test-blog"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("example"),
		postgres.WithInitScripts(filepath.Join(".", "testdata", "init-db.sql")),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		return nil, err
	}

	connStr, err := pgContainer.ConnectionString(ctx,
		"sslmode=disable",
		"timezone=UTC",
	)
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}, nil
}
