package repositories_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gera9/blog/internal/repositories/testhelpers"
	"github.com/gera9/blog/pkg/postgres"
)

var (
	PostgresContainer *testhelpers.PostgresContainer
	PostgresConn      *postgres.Postgres
)

func TestMain(m *testing.M) {
	SetUp()
	code := m.Run()
	TearDown()
	os.Exit(code)
}

func SetUp() {
	ctx := context.Background()
	var err error

	PostgresContainer, err = testhelpers.NewPostgresContainer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = PostgresContainer.Snapshot(ctx)
	if err != nil {
		log.Fatal(err)
	}

	PostgresConn, err = postgres.NewPostgres(ctx, PostgresContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func TearDown() {
	PostgresConn.Close()
	ctx := context.Background()
	if err := PostgresContainer.Terminate(ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}
