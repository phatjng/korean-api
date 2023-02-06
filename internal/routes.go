package internal

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phatjng/korean-api/db/sqlite"
	"github.com/phatjng/korean-api/internal/utils"
)

type Router struct {
	ctx     context.Context
	mux     *chi.Mux
	queries *sqlite.Queries
}

func NewRouter(queries *sqlite.Queries) *Router {
	return &Router{
		ctx:     context.Background(),
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
		// "/api/card" - Card routes
		r.Route("/card", func(r chi.Router) {
			r.Post("/", rt.CreateCard)
			r.Get("/", rt.GetCard)
		})

		// "/api/deck" - Deck routes
		r.Route("/deck", func(r chi.Router) {
			r.Post("/", rt.CreateDeck)
			r.Get("/", rt.GetDeck)
		})

	})

	return rt.mux
}

// Create deck
func (rt *Router) CreateCard(w http.ResponseWriter, r *http.Request) {
	var body Card

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Card body cannot be empty
	if len(body.Front) == 0 || len(body.Back) == 0 || len(body.DeckTitle) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get deck by title
	deck, err := rt.queries.GetDeckByTitle(rt.ctx, body.DeckTitle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert card to database
	id := utils.GenerateID()

	err = rt.queries.CreateCard(rt.ctx, sqlite.CreateCardParams{
		ID:     id,
		Front:  body.Front,
		Back:   body.Back,
		DeckID: deck.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Get card
func (rt *Router) GetCard(w http.ResponseWriter, r *http.Request) {
	var body Card

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	card, err := rt.queries.GetCard(rt.ctx, body.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(card)
}

// Create deck
func (rt *Router) CreateDeck(w http.ResponseWriter, r *http.Request) {
	var body Deck

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Deck body cannot be empty
	if len(body.Title) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert deck to database
	id := utils.GenerateID()

	err = rt.queries.CreateDeck(rt.ctx, sqlite.CreateDeckParams{
		ID:    id,
		Title: body.Title,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Get deck. Returns all the cards from the deck
func (rt *Router) GetDeck(w http.ResponseWriter, r *http.Request) {
	var body Deck

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var deck sqlite.Deck

	// If title is not provided, get deck by id
	if len(body.Title) == 0 {
		deck, err = rt.queries.GetDeck(rt.ctx, body.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// If id is not provided, get deck by title
	if len(body.ID) == 0 {
		deck, err = rt.queries.GetDeckByTitle(rt.ctx, body.Title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	cards, err := rt.queries.GetAllCardsFromDeck(rt.ctx, deck.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := &Deck{
		ID:    deck.ID,
		Title: deck.Title,
		Cards: cards,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
