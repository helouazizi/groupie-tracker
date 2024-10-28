// handlers/handlers.go
package handlers

import (
	"groupie-tracker/api"
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
	}
	if err := api.FetchArtists(); err != nil {
		log.Fatal(err)
	}
	tmple, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmple.Execute(w, api.Artists)
	if err != nil {
		log.Fatal(err)
	}

}

/*
lets serve our static folder for any assets request
*/
func ServStatic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static/" || r.URL.Path == "/static" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.StripPrefix("/static", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)

}
