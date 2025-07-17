package main

import (
	"log"

	"github.com/joho/godotenv"
)

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}
}
