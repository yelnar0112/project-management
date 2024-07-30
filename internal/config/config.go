package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	} else {
		log.Println("Configuration loaded from .env file")
	}
}
