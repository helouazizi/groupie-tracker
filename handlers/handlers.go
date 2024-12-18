// handlers/handlers.go
package handlers

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"groupie-tracker/api"
)

type Pages struct {
	homePage   *template.Template
	Artistpage *template.Template
	ErrorPage  *template.Template
}

var pages Pages

func init() {
	var err error
	pages.homePage, err = template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	pages.Artistpage, err = template.ParseFiles("templates/artist.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	pages.ErrorPage, err = template.ParseFiles("templates/error.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		pages.ErrorPage.Execute(w, "NOT FOUND")
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		pages.ErrorPage.Execute(w, "METHOD NOT ALLOWED")
		return
	}

	w.Header().Set("Content-Type", "text/html")
	artists := api.Artists

	if err := pages.homePage.Execute(w, artists); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		pages.ErrorPage.Execute(w, "Internal server error")
		return
	}
}

/*
lets serve our static folder for any assets request
*/
func ServStatic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static/" || r.URL.Path == "/static/css" || r.URL.Path == "/static/css/" {
		w.WriteHeader(http.StatusNotFound)
		pages.ErrorPage.Execute(w, "NOT FOUND")
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		pages.ErrorPage.Execute(w, "METHOD NOT ALLOWED")
		return
	}
	_, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		pages.ErrorPage.Execute(w, "NOT FOUND")
		return
	}
	http.StripPrefix("/static", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
}

func ArtistDetails(w http.ResponseWriter, r *http.Request) {
	// get the artist id from the url
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		pages.ErrorPage.Execute(w, "METHOD NOT ALLOWED")
		return
	}
	w.Header().Set("Content-Type", "text/html")
	id := r.URL.Query().Get("id")
	artist_details, err := api.FetchDetails(id) // Get the artist details
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		pages.ErrorPage.Execute(w, "BAD REQUEST")
		return
	}

	if err := pages.Artistpage.Execute(w, artist_details); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		pages.ErrorPage.Execute(w, "Internal server error")
		return
	}
}

/*
this funtion is special for the errors when occurs
*/
