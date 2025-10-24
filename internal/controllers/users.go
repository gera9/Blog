package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/google/uuid"

	"github.com/gera9/blog/internal/controllers/dtos"
	"github.com/gera9/blog/internal/models"
	"github.com/gera9/blog/pkg/middlewares"
	"github.com/go-chi/chi/v5"
)

type UsersService interface {
	CreateUser(ctx context.Context, user models.User) (uuid.UUID, error)
	FindAllUsers(ctx context.Context, limit, offset int) ([]models.User, error)
	FindUserById(ctx context.Context, id uuid.UUID) (models.User, error)
	UpdateUserById(ctx context.Context, id uuid.UUID, user models.User) error
	DeleteUserById(ctx context.Context, id uuid.UUID) error
}

type usersController struct {
	usersService UsersService
}

func NewUsersController(usersService UsersService) *usersController {
	return &usersController{usersService}
}

func (c usersController) Routes(mm *middlewares.MiddlewareManager) *chi.Mux {
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

func (c usersController) Create(w http.ResponseWriter, r *http.Request) {
	userPayload := dtos.CreateUser{}
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	id, err := c.usersService.CreateUser(r.Context(), userPayload.ToUser())
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

func (c usersController) FindAll(w http.ResponseWriter, r *http.Request) {
	limit := r.Context().Value(middlewares.ContextKeyLimit).(int)
	offset := r.Context().Value(middlewares.ContextKeyOffset).(int)

	users, err := c.usersService.FindAllUsers(r.Context(), limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	response := make([]dtos.UserResponse, len(users))
	for i, user := range users {
		response[i] = toUserResponse(user)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c usersController) FindById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
		return
	}

	user, err := c.usersService.FindUserById(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(toUserResponse(user))
}

func (c usersController) UpdateById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
		return
	}

	userPayload := dtos.UpdateUser{}
	if err := json.NewDecoder(r.Body).Decode(&userPayload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = c.usersService.UpdateUserById(r.Context(), id, userPayload.ToUser())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c usersController) DeleteById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
		return
	}

	err = c.usersService.DeleteUserById(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func toUserResponse(user models.User) dtos.UserResponse {
	return dtos.UserResponse(user)
}
