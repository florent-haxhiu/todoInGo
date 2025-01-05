package router

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"florent-haxhiu/todoInGo/internal/model"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Route("/note", func(r chi.Router) {
		r.Use(authorizeSession)
		r.Get("/", GetAllNotes)
		r.Post("/", PostNote)
		r.Route("/{noteId}", func(r chi.Router) {
			r.Get("/", GetNote)
			r.Put("/", UpdateNote)
		})
	})
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", Register)
		r.Post("/login", Login)
	})

	return r
}

func authorizeSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		claims, err := verifyToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		k := model.UserId("userId")

		ctx := context.WithValue(r.Context(), k, claims.UserId.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
