package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gera9/blog/internal/controllers"
	"github.com/gera9/blog/internal/repositories"
	"github.com/gera9/blog/internal/services"
	"github.com/gera9/blog/pkg/middlewares"
)

func main() {

	mongoUri := os.Getenv("MONGO_URL")
	databaseName := os.Getenv("MONGO_DATABASE_NAME")
	r, err := repositories.NewRepositories(context.Background(), mongoUri, databaseName)
	if err != nil {
		log.Fatal(err)
	}

	s := &services.Services{UsersRepository: r, PostsRepository: r}
	mm := &middlewares.MiddlewareManager{}

	http.ListenAndServe(":3000", controllers.BuildRoutes(s, mm))
}
