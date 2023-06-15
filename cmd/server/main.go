package main

import (
	"github.com/porky256/mock-interview-tbr/internal/routes"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
