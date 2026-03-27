package main

import (
	"log"

	"go-search-api/internal/server"
)

func main() {
	srv := server.New(":8080")

	log.Printf("listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
