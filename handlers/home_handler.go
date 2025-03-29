package handlers

import (
	"encoding/json"
	"groupie-tracker/repository"
	"net/http"
)

type Home_handler struct {
	Store *repository.Store
}

func (h *Home_handler) Home(w http.ResponseWriter, r *http.Request) {
	all_artist := h.Store.GetArtists()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(all_artist)
}
