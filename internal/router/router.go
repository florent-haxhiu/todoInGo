package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Route("/note", func(r chi.Router) {
		r.Get("/", GetAllNotes)
		r.Post("/", PostNote)
		r.Route("/{noteId}", func(r chi.Router) {
			r.Get("/", GetNote)
			r.Put("/", UpdateNote)
		})
	})
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", Register)
		r.Post("/login", Login)
	})

	return r
}
