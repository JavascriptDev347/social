package main

import (
	"log"
	"net/http"
	"time"

	"github.com/JavascriptDev347/social.git/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr     string
	dbConfig dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))
	// r.Group()
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthHandler)

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.CreatePostHandler)
			r.Route("/{postID}", func(r chi.Router) {
				r.Get("/", app.GetPostHandler)
			})
		})
	})

	return r

}

func (app *application) run(mux *chi.Mux) error {

	src := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server has started at %s", app.config.addr)
	return src.ListenAndServe()
}
