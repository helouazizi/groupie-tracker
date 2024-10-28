// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/api"
	"groupie-tracker/handlers"
)

func init() {
	if err := api.FetchArtists(); err != nil {
		log.Fatalf("Failed to fetch artists: %v", err)
	}

	// Repeat for locations, dates, and relation
}

func main() {
	// fmt.Println("Artists:", api.Artists)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/static/", handlers.ServStatic)
	// Output other data as needed
	fmt.Println("server listnign on port  8080 >> http://localhost:8080 ")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
