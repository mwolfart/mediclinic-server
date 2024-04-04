package setup

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func StartDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	log.Println("DB connected")
	return db, nil
}
