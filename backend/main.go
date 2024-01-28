package main

import (
	"net/http"
	"strings"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".") {
		http.ServeFile(w, r, "../frontend/build"+r.URL.Path)
		return
	}

	if r.URL.Path == "/" {
		serveReactApp(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func serveReactApp(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("../frontend/build"))
	http.StripPrefix("/", fs).ServeHTTP(w, r)
}

func main() {
	http.Handle("/", &myHandler{})
	http.ListenAndServe(":8080", nil)
}
