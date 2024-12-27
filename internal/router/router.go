package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"florent-haxhiu/todoInGo/internal/controllers"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Route("/note", func(r chi.Router) {
		r.Get("/", controllers.GetAllNotes)
		r.Post("/", controllers.PostNote)
		r.Route("/{noteId}", func(r chi.Router) {
			r.Get("/", controllers.GetNote)
			r.Put("/", controllers.UpdateNote)
		})
	})

	return r
}
