package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Cesar90/golang-social-api/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
}

// func (app *application) mount() *http.ServeMux {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

// 	return mux
// }

// func (app *application) mount() *chi.Mux {
/* Here we are use http.Handler interface because *chi.Mux implement it*/
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// Set a timeout value on the request context (ctx), that will signal
	// throuh ctx.Done() that the request has timed out and futher
	// processing should be stopped
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})
	return r
}

// func (app *application) run(mux *chi.Mux) error {
func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started %s", app.config.addr)

	return srv.ListenAndServe()
}
