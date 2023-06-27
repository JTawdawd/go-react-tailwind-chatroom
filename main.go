package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)
	s := &server{}
	http.Handle("/", corsMiddleware(s))
	http.HandleFunc("/login", HandleLogin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
