package handlers

import (
	"encoding/json"
	"groupie-tracker/api"
	"groupie-tracker/models"
	"groupie-tracker/repository"
	"log"
	"net/http"
	"sync"
)

type ArtistDeatils struct {
	Store *repository.Store
}

func (h *ArtistDeatils) ArtistDetail(w http.ResponseWriter, r *http.Request) {
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

	// Define variables to hold fetched data
	var locationData, concertData, relationData interface{}
	var wg sync.WaitGroup
	errChan := make(chan error, 3)

	// Function to fetch data
	fetchData := func(url string, target interface{}) {
		defer wg.Done()
		if err := api.Fetch(url, target); err != nil {
			log.Println("Error fetching data:", err)
			errChan <- err
		}
	}

	// Concurrently fetch data
	wg.Add(3)
	go fetchData(artist.Locations, &locationData)
	go fetchData(artist.ConcertDates, &concertData)
	go fetchData(artist.Relations, &relationData)

	// Wait for all fetch operations to complete
	wg.Wait()
	close(errChan)

	// Check if any error occurred
	if len(errChan) > 0 {
		http.Error(w, "Failed to fetch artist details", http.StatusInternalServerError)
		return
	}

	// Combine artist and fetched data into a response
	extendedArtist := struct {
		Artist   models.Artist `json:"artist"`
		Location interface{}   `json:"locationData"`
		Concert  interface{}   `json:"concertData"`
		Relation interface{}   `json:"relationData"`
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
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
