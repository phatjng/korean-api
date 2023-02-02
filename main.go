package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/phatjng/korean-api/db/sqlite"
	"github.com/phatjng/korean-api/internal"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/database.sqlite3")
	if err != nil {
		panic(err)
	}

	queries := sqlite.New(db)

	r := internal.NewRouter(queries)

	fmt.Println("🚀 Running...")

	if err := http.ListenAndServe(":8000", r.Register()); err != nil {
		log.Fatal(err)
	}
}
