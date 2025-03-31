package router

import (
	"groupie-tracker/handlers"
	"groupie-tracker/repository"
	"net/http"
)

func NewRouter(s *repository.Store) *http.ServeMux {
	mux := http.NewServeMux()
	all := &handlers.Home_handler{Store: s}
	//deatils := &handlers.ArtistDeatils{Store: s}
	mux.HandleFunc("/", all.Home)
	//mux.HandleFunc("/artist", deatils.ArtistDetail)
	return mux
}
