package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/repository"
	"log"
	"net/http"
)

type Artist_Deatils struct {
	Store *repository.Store
}

func (h *Artist_Deatils) Artist_Deatil(w http.ResponseWriter, r *http.Request) {
	//	// lets get the id
	//	id := r.Header.Get("id")
	//
	path := r.URL
	fmt.Println(path)
	// Extract ID from query parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing artist ID", http.StatusBadRequest)
		return
	}

	// Find artist in store
	artist, found := h.Store.GetArtistByID(id)
	if !found {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(artist); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
