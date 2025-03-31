package handlers

import (
	"encoding/json"
	"groupie-tracker/models"
	"groupie-tracker/repository"
	"net/http"
)

type ArtistDeatils struct {
	Store *repository.Store
}

func (h *ArtistDeatils) ArtistDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Fix CORS issues
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
	locations, exist := h.Store.GetLocationById(id)
	if !exist {
		http.Error(w, "Locations not found", http.StatusNotFound)
		return
	}
	relations, exist := h.Store.GetRealtionById(id)
	if !exist {
		http.Error(w, "Realtions not found", http.StatusNotFound)
		return
	}
	dates, exist := h.Store.GetDateById(id)
	if !exist {
		http.Error(w, "Locations not found", http.StatusNotFound)
		return
	}

	// Function to fetch data
	// fetchData := func(url string, target interface{}) {
	// 	defer wg.Done()
	// 	if err := api.Fetch(url, target); err != nil {
	// 		log.Println("Error fetching data:", err)
	// 		errChan <- err
	// 	}
	// }

	// Concurrently fetch data
	//wg.Add(3)
	// go fetchData(artist.Locations, &locationData)
	// go fetchData(artist.ConcertDates, &concertData)
	// go fetchData(artist.Relations, &relationData)

	// Combine artist and fetched data into a response
	extendedArtist := struct {
		Artist   models.Artist
		Locations models.Location
		Dates    models.Date
		Relations models.Relation
	}{
		Artist:   artist,
		Locations: locations,
		Dates:    dates,
		Relations: relations,
	}

	// Respond with the extended artist JSON data

	if err := json.NewEncoder(w).Encode(extendedArtist); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
