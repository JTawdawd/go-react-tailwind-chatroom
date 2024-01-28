package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func handlePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	response, err := handlerDefinitions[r.URL.Path](body)
	log.Print(err)
	if err != nil {
		response, _ = json.Marshal(map[string]interface{}{
			"status": "Error",
			"Error":  err,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	var response []byte

	response, err := handlerDefinitions[r.URL.Path]([]byte{})

	if err != nil {
		response, _ = json.Marshal(map[string]interface{}{
			"status": "Error",
			"Error":  err,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
