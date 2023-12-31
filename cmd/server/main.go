package main

import (
	"log"
	"net/http"
	"time"

	"github.com/porky256/mock-interview-tbr/internal/routes"
)

func main() {
	log.Printf("Server started")

	router := routes.NewRouter()

	const timeout = 3 * time.Second

	server := &http.Server{
		Addr:              "localhost:8080",
		Handler:           router,
		ReadHeaderTimeout: timeout,
	}
	log.Fatal(server.ListenAndServe())
}
