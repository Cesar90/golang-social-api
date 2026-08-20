package main

import (
	"log"

	"github.com/Cesar90/golang-social-api/internal/env"
	"github.com/Cesar90/golang-social-api/internal/store"
)

func main() {
	cfg := config{
		// addr: ":8080",
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
