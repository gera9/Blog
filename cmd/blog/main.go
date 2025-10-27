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
	"github.com/gera9/blog/pkg/utils"
)

func main() {
	postgresConnStr := os.Getenv("POSTGRES_URL")
	r, err := repositories.NewRepositories(context.Background(), postgresConnStr, utils.RealClock{})
	if err != nil {
		log.Fatal(err)
	}

	s := &services.Services{UsersRepository: r, PostsRepository: r}
	mm := &middlewares.MiddlewareManager{}

	addr := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	log.Println("Listening on addr:", addr)

	http.ListenAndServe(addr, controllers.BuildRoutes(s, mm))
}
