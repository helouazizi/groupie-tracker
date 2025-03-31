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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Fix CORS issues
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	all_artist := h.Store.GetArtists()
	if err := json.NewEncoder(w).Encode(all_artist); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
