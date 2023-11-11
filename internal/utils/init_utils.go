package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Environment variables were not initialized: %s", err)
	}
}
