package main

import (
	"log"

	"github.com/JavascriptDev347/social.git/internal/db"
	"github.com/JavascriptDev347/social.git/internal/env"
	"github.com/JavascriptDev347/social.git/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	store := store.NewStorage(conn)
	db.Seed(store)
}
