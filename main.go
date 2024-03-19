package main

import (
	"log"
	"medclin/handlers"
	"medclin/setup"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := setup.StartDB()
	if err != nil {
		log.Fatal("Error opening connection to database")
		os.Exit(1)
	}

	handlers.SetupHandlers(db)
	setup.StartWS()
}
