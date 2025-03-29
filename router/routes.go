package router

import (
	"groupie-tracker/handlers"
	"groupie-tracker/repository"
	"net/http"
)

func NewRouter(s *repository.Store) *http.ServeMux {
	mux := http.NewServeMux()
	toHandle := &handlers.Home_handler{Store: s}
	mux.HandleFunc("/", toHandle.Home)
	return mux
}
