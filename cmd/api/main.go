package main

import (
	"log"

	"github.com/Cesar90/golang-social-api/internal/env"
)

func main() {
	cfg := config{
		// addr: ":8080",
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
