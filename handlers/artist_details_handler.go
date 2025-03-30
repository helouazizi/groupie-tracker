package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/api"
	"groupie-tracker/models"
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
	// Fetch additional data from the provided API URLs
	var locationData, concertData, relationData interface{}

	// Fetch locations data
	if err := api.Fetch(artist.Locations, &locationData); err != nil {
		log.Println("Error fetching locations:", err)
	}
	// Fetch concert dates data
	if err := api.Fetch(artist.ConcertDates, &concertData); err != nil {
		log.Println("Error fetching concert dates:", err)
	}
	// Fetch relations data
	if err := api.Fetch(artist.Relations, &relationData); err != nil {
		log.Println("Error fetching relations:", err)
	}

	// Extend artist data with the fetched data
	extendedArtist := struct {
		Artist   models.Artist `json:"artist"`
		Location interface{}    `json:"locationData"`
		Concert  interface{}    `json:"concertData"`
		Relation interface{}    `json:"relationData"`
	}{
		Artist:   artist,
		Location: locationData,
		Concert:  concertData,
		Relation: relationData,
	}

	// Respond with the extended artist JSON data
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Fix CORS issues

	if err := json.NewEncoder(w).Encode(extendedArtist); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

}
