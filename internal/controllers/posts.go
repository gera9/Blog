package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gera9/blog/internal/controllers/dtos"
	"github.com/gera9/blog/internal/models"
	"github.com/gera9/blog/pkg/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type PostsService interface {
	CreatePost(ctx context.Context, post models.Post) (uuid.UUID, error)
	FindAllPosts(ctx context.Context, limit, offset int, authorId uuid.UUID) ([]models.Post, error)
	FindPostByIdAndAuthorId(ctx context.Context, id, authorId uuid.UUID) (models.Post, error)
	UpdatePostByIdAndAuthorId(ctx context.Context, id, authorId uuid.UUID, post models.Post) error
	DeletePostById(ctx context.Context, id uuid.UUID) error
}

type postsController struct {
	postsService PostsService
}

func NewPostsController(postsService PostsService) *postsController {
	return &postsController{postsService}
}

func (c postsController) Routes(mm *middlewares.MiddlewareManager) *chi.Mux {
	r := chi.NewMux()

	r.Post("/", c.Create)
	r.With(mm.List).Get("/", c.FindAll)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", c.FindById)
		r.Patch("/", c.UpdateById)
		r.Delete("/", c.DeleteById)
	})

	return r
}

func (c postsController) Create(w http.ResponseWriter, r *http.Request) {
	postPayload := dtos.CreatePost{}
	err := json.NewDecoder(r.Body).Decode(&postPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	id, err := c.postsService.CreatePost(r.Context(), postPayload.ToPost())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"id": id,
	})
}

func (c postsController) FindAll(w http.ResponseWriter, r *http.Request) {
	limit := r.Context().Value(middlewares.ContextKeyLimit).(int)
	offset := r.Context().Value(middlewares.ContextKeyOffset).(int)

	q := r.URL.Query()
	authorIdStr := q.Get("author_id")
	authorId, err := uuid.Parse(authorIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid author_id UUID format",
		})
		return
	}

	posts, err := c.postsService.FindAllPosts(r.Context(), limit, offset, authorId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	response := make([]dtos.PostResponse, len(posts))
	for i, post := range posts {
		response[i] = toPostResponse(post)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c postsController) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid UUID format",
		})
		return
	}

	q := r.URL.Query()
	authorIdStr := q.Get("author_id")
	authorId, err := uuid.Parse(authorIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid author_id UUID format",
		})
		return
	}

	post, err := c.postsService.FindPostByIdAndAuthorId(r.Context(), id, authorId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(toPostResponse(post))
}

func (c postsController) UpdateById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid UUID format",
		})
		return
	}

	q := r.URL.Query()
	authorIdStr := q.Get("author_id")
	authorId, err := uuid.Parse(authorIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid author_id UUID format",
		})
		return
	}

	postPayload := dtos.CreatePost{}
	err = json.NewDecoder(r.Body).Decode(&postPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = c.postsService.UpdatePostByIdAndAuthorId(r.Context(), id, authorId, postPayload.ToPost())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c postsController) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid UUID format",
		})
		return
	}

	err = c.postsService.DeletePostById(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func toPostResponse(post models.Post) dtos.PostResponse {
	return dtos.PostResponse(post)
}
