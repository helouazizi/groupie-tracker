package main

import (
	"groupie-tracker/repository"
	"groupie-tracker/router"
	"log"
	"net/http"
)

func main() {
	store := repository.New_Store()
	store.LoadData()
	mux := router.NewRouter(store)
	log.Println("Server is running on http://localhost:8080/")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
