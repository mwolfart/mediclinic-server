package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

const PORT = 8080

type Person struct {
	Name    string
	Address string
	Age     int
}

func handlePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var p Person

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Bad request body.", http.StatusBadRequest)
		return
	}

	fmt.Printf("%s\n", p.Name)
	fmt.Fprintf(w, "POST successful!\n")
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Printf("%s, %s\n", name, address)
	fmt.Fprintf(w, "%s, %s\n", name, address)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func setupDB() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	checkError(err)

	defer db.Close()

	err = db.Ping()
	checkError(err)

	fmt.Println("DB connected")
	return db
}

func main() {
	fmt.Printf("Starting server at port %d\n", PORT)
	db := setupDB()
	// setupDB()

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", PORT),
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Println("Shutting down server gracefully...")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	// gin.Gin()
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/person", handlePerson)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	updateStmt := `insert into "Test" values "FOO"=abc, "BAR"=xyz`
	_, e := db.Exec(updateStmt)
	checkError(e)
}
