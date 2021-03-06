package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvKey(key string) string {

	// load .env file
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
