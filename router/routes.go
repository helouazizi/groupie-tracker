package router

import (
	"groupie-tracker/handlers"
	"groupie-tracker/repository"
	"net/http"
)

func NewRouter(s *repository.Store) *http.ServeMux {
	mux := http.NewServeMux()
	all := &handlers.HomeHandler{Store: s}
	deatils := &handlers.ArtistDeatils{Store: s}
	filter := &handlers.FilterHandler{Store: s}
	mux.HandleFunc("/", all.Home)
	mux.HandleFunc("/artist", deatils.ArtistDetail)
	mux.HandleFunc("/filter", filter.Filter)
	return mux
}
