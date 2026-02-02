package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gera9/blog/internal/controllers"
	"github.com/gera9/blog/internal/repositories"
	"github.com/gera9/blog/internal/services"
	"github.com/gera9/blog/pkg/middlewares"
	"github.com/gera9/blog/pkg/postgres"
	"github.com/gera9/blog/pkg/utils"
)

func main() {
	postgresConnStr := os.Getenv("POSTGRES_URL")
	postgresConn, err := postgres.NewPostgres(context.Background(), postgresConnStr)
	if err != nil {
		log.Fatal(err)
	}

	postsRepo := repositories.NewPostsRepository(postgresConn, utils.RealClock{})
	usersRepo := repositories.NewUsersRepository(postgresConn, utils.RealClock{})

	postsServ := services.NewPostsService(postsRepo)
	usersServ := services.NewUsersService(usersRepo)

	mm := &middlewares.MiddlewareManager{}

	addr := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	log.Println("Listening on addr:", addr)

	http.ListenAndServe(addr, controllers.BuildRoutes(mm, usersServ, postsServ))
}
