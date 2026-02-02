package repositories_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gera9/blog/internal/repositories"
	"github.com/gera9/blog/internal/repositories/testhelpers"
	"github.com/gera9/blog/pkg/utils"
)

var (
	PostgresContainer *testhelpers.PostgresContainer
	Repository        *repositories.Repositories
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

	Repository, err = repositories.NewRepositories(ctx, PostgresContainer.ConnectionString, utils.MockClock{})
	if err != nil {
		log.Fatal(err)
	}

}

func TearDown() {
	Repository.Close()
	ctx := context.Background()
	if err := PostgresContainer.Terminate(ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}
