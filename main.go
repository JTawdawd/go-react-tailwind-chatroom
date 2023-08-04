package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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

func connectDB() {
	pgsqlDetails := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, os.Getenv("DB_PASSWORD"), dbname)
	db, err := sql.Open("postgres", pgsqlDetails)
	if err != nil {
		panic(err)
	}
	DB = db
	log.Println("Connected to database :)")
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {

	fs := http.FileServer(http.Dir("frontend/build"))
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	connectDB()

	http.Handle("/", fs)
	http.HandleFunc("/login", HandleLogin)
	http.HandleFunc("/create", CreateUser)
	http.HandleFunc("/chatrooms", Chatrooms)
	http.HandleFunc("/chatroom", GetChatroom)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
