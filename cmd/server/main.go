package main

import (
	"github.com/porky256/mock-interview-tbr/internal/routes"

	"log"
	"net/http"
	"time"
)

func main() {
	log.Printf("Server started")

	router := routes.NewRouter()

	const timeout = 3 * time.Second

	server := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: timeout,
	}
	log.Fatal(server.ListenAndServe())
}
