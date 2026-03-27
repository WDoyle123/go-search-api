package server

import (
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /", root)
	mux.HandleFunc("GET /search", search)
	return logging(mux)
}

