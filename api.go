package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	// Enable CORS
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	// Set up routes and handlers
	http.HandleFunc("/login", HandleLogin)

	// Start the server with CORS middleware
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(nil)))
}
