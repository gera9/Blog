package middlewares

import (
	"context"
	"net/http"
	"strconv"
)

var (
	ContextKeyLimit  = &ContextKey{"limit"}
	ContextKeyOffset = &ContextKey{"offset"}
)

func (mm MiddlewareManager) List(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		limitStr := q.Get("limit")
		if limitStr == "" {
			limitStr = "10"
		}

		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "invalid limit value", http.StatusBadRequest)
			return
		}

		offsetStr := q.Get("offset")
		if offsetStr == "" {
			offsetStr = "0"
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			http.Error(w, "invalid offset value", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, ContextKeyLimit, limit)
		ctx = context.WithValue(ctx, ContextKeyOffset, offset)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
