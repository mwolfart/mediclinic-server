package setup

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func StartDB() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("Error opening connection to database")
		return nil
	}

	defer db.Close()

	log.Println("DB connected")
	return db
}
