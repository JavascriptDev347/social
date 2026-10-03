package main

import (
	"log"
	"os"

	"github.com/JavascriptDev347/social.git/internal/env"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}

	os.LookupEnv("PATH")

	mux := app.mount()
	log.Fatal(app.run(mux))

}
