package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "postgres"
)

var DB *sql.DB

type server struct{}

func connectDB() {
	pgsqlDetails := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, os.Getenv("DB_PASSWORD"), dbname)
	db, err := sql.Open("postgres", pgsqlDetails)
	if err != nil {
		panic(err)
	}
	DB = db
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	connectDB()
	getUser1()

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)
	s := &server{}
	http.Handle("/", corsMiddleware(s))
	http.HandleFunc("/login", HandleLogin)
	http.HandleFunc("/create", CreateUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
