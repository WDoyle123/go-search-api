package server

import (
	"net/http"
	"time"
)

func New(addr string) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}
}
