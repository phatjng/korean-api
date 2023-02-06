package internal

import "github.com/phatjng/korean-api/db/sqlite"

type Card struct {
	ID        string `json:"id"`
	Front     string `json:"front"`
	Back      string `json:"back"`
	DeckTitle string `json:"deck_title"`
}

type Deck struct {
	ID    string        `json:"id"`
	Title string        `json:"title"`
	Cards []sqlite.Card `json:"cards"`
}
