// handlers/handlers.go
package handlers

import (
	"fmt"
	"groupie-tracker/api"
	"groupie-tracker/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
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
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

/*
lets serve our static folder for any assets request
*/
func ServStatic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static/" || r.URL.Path == "/static" {
		w.WriteHeader(http.StatusNotFound)
		pages.ErrorPage.Execute(w, "NOT FOUND")
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		pages.ErrorPage.Execute(w, "METHOD NOT ALLOWED")
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
    artist, err := getArtistByID(id) // Get the artist details
    if err != nil {
        http.Error(w, "Artist not found", http.StatusNotFound)
        return
    }

	if err := pages.Artistpage.Execute(w, artist); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func getArtistByID(id string) (*models.Artist, error) {
	ids,_ := strconv.Atoi(id)
	for _, artist := range *api.Artists { // 'artists' is a slice of Artist structs
		if artist.ID == ids {
			return &artist, nil
		}
	}
	return nil, fmt.Errorf("artist not found")
}

/*
this funtion is special for the errors when occurs
*/
