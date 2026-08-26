package main

import (
	"log"

	"github.com/Cesar90/golang-social-api/internal/db"
	"github.com/Cesar90/golang-social-api/internal/env"
	"github.com/Cesar90/golang-social-api/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5432/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)

}
