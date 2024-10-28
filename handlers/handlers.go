package handlers

import (
	"html/template"
	"log"
	"net/http"

	"groupie-tracker/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
	}
	tmple, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}
	artist := models.Artist{
		Name:  "John Lennon",
		Image: "",
	}

	tmple.Execute(w, artist)
}
