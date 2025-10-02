package controllers

import (
	"net/http"

	"github.com/gera9/blog/internal/services"
	"github.com/gera9/blog/pkg/middlewares"
	"github.com/go-chi/chi/v5"
)

func BuildRoutes(s *services.Services, mm *middlewares.MiddlewareManager) http.Handler {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", NewUsersController(s).Routes(mm))
		r.Mount("/posts", NewPostsController(s).Routes(mm))
	})

	return r
}
