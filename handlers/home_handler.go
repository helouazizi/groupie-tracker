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
	w.Header().Set("Access-Control-Allow-Origin", "*") // Fix CORS issues
	//artists := h.Store.GetArtists() // Ensure this returns data
	if err := json.NewEncoder(w).Encode(all_artist); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
