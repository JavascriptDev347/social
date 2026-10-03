package main

import (
	"log"

	"github.com/JavascriptDev347/social.git/internal/env"
	"github.com/JavascriptDev347/social.git/internal/store"
)

func main() {

	cfg := config{
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
