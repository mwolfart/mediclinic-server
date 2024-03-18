package main

import (
	"log"
	"medclin/setup"
	"os"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	db, err := setup.StartDB()
	if err != nil {
		log.Fatal("Error opening connection to database")
		os.Exit(1)
	}
	boil.SetDB(db)
	setup.StartWS()
}
