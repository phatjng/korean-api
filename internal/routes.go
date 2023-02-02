package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phatjng/korean-api/db/sqlite"
)

type Router struct {
	mux     *chi.Mux
	queries *sqlite.Queries
}

func NewRouter(queries *sqlite.Queries) *Router {
	return &Router{
		mux:     chi.NewRouter(),
		queries: queries,
	}
}

func (rt *Router) Register() *chi.Mux {
	rt.mux.Use(middleware.RequestID)
	rt.mux.Use(middleware.RealIP)
	rt.mux.Use(middleware.Logger)
	rt.mux.Use(middleware.Recoverer)

	rt.mux.Route("/api", func(r chi.Router) {
		// "/api/deck" - Deck routes
		r.Route("/deck", func(r chi.Router) {
			r.Get("/", rt.GetDeck)
		})

		// "/api/card" - Card routes
		r.Route("/card", func(r chi.Router) {
			r.Get("/", rt.GetCard)
		})
	})

	return rt.mux
}

// Get deck
func (rt *Router) GetDeck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET DECK"))
}

// Get card
func (rt *Router) GetCard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET CARD"))
}
