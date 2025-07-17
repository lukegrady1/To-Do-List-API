package main

import (
	"log"
	"os"

	"github.com/lukegrady1/to-do-list-api/handlers"
	"github.com/lukegrady1/to-do-list-api/router"
)

func main() {
	initConfig()
	initDB()
	handlers.DB = DB

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := router.SetupRouter()
	log.Printf("Starting server on :%s…\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
