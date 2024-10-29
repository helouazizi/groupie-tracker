// handlers/handlers.go
package handlers

import (
	"groupie-tracker/api"
	"log"
	"net/http"
	"text/template"
)

type Pages struct {
	homePage *template.Template
	ErrorPage  *template.Template

}

var pages Pages

func init() {
	var err error
	pages.homePage, err = template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.StripPrefix("/static", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)

}
