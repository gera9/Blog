package controllers

import (
	"net/http"

	"github.com/gera9/blog/pkg/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func BuildRoutes(mm *middlewares.MiddlewareManager, usersService UsersService, postsService PostsService) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", NewUsersController(usersService).Routes(mm))
		r.Mount("/posts", NewPostsController(postsService).Routes(mm))
	})

	return r
}
