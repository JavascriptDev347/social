package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	// Loads .env from the current working directory.
	// Error is ignored on purpose: in production the file won't exist
	// and real environment variables are used instead.
	_ = godotenv.Load()
}

func GetString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	valHasInt, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return valHasInt
}
