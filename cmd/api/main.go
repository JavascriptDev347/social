package main

import (
	"fmt"
	"log"

	"github.com/JavascriptDev347/social.git/internal/db"
	"github.com/JavascriptDev347/social.git/internal/env"
	"github.com/JavascriptDev347/social.git/internal/store"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		dbConfig: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.dbConfig.addr, cfg.dbConfig.maxOpenConns, cfg.dbConfig.maxIdleConns, cfg.dbConfig.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Database connection establish")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
