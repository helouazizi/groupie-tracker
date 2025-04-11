package handlers

import (
	"encoding/json"
	"groupie-tracker/models"
	"groupie-tracker/repository"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type FilterRequest struct {
	CreationFrom string `json:"creationDateFrom"`
	CreationTo   string `json:"creationDateTo"`
	AlbumFrom    string `json:"firstAlbumFrom"`
	AlbumTo      string `json:"firstAlbumTo"`
	Members      string `json:"members"`
	ConcertDate  string `json:"concertDates"`
}

type FilterHandler struct {
	Store *repository.Store
}

func (f *FilterHandler) Filter(w http.ResponseWriter, r *http.Request) {
	// ✅ CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// ✅ Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var filterReq FilterRequest
	if err := json.NewDecoder(r.Body).Decode(&filterReq); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	var filtered []models.Artist
	f.applyFilters(filterReq, &filtered)

	if len(filtered) == 0 {
		http.Error(w, "No matching artists found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(filtered); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}

// applyFilters performs AND logic across all conditions
func (f *FilterHandler) applyFilters(data FilterRequest, result *[]models.Artist) {
	// Parse filter values
	creationFrom, err1 := strconv.Atoi(data.CreationFrom)
	creationTo, err2 := strconv.Atoi(data.CreationTo)
	albumFrom, err3 := strconv.Atoi(data.AlbumFrom)
	albumTo, err4 := strconv.Atoi(data.AlbumTo)
	members, err5 := strconv.Atoi(data.Members)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		log.Println("Invalid filter input")
	}

	// Normalize the concert query
	concerts := strings.ToLower(data.ConcertDate)

	// Loop through all artists
	for _, artist := range f.Store.Artists {
		// ✅ Filter by creation date (after 2010)
		if creationFrom != 0 && artist.CreationDate <= creationFrom {
			continue
		}
		if creationTo != 0 && artist.CreationDate > creationTo {
			continue
		}

		// ✅ Filter by first album year (after 2010)
		parts := strings.Split(artist.FirstAlbum, "-")
		if len(parts) != 3 {
			continue
		}
		albumYear, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}
		if albumFrom != 0 && albumYear <= albumFrom {
			continue
		}
		if albumTo != 0 && albumYear > albumTo {
			continue
		}
		if members != 0 && (len(artist.Members) != members || members <= len(artist.Members)) {
			continue
		}
		if concerts != "" && !f.matchLocation(artist.ID, concerts) {
			continue
		}

		// Add the artist to the filtered list if all conditions match
		*result = append(*result, artist)
	}
}

// matchLocation checks if the artist has a concert in the given location
// matchLocation checks if the artist has a concert in the given location
func (f *FilterHandler) matchLocation(id int, query string) bool {
	// Ensure the query is case-insensitive and formatted consistently
	query = strings.ToLower(strings.TrimSpace(query))
	// Replace commas with hyphens for better matching
	query = strings.ReplaceAll(query, ",", "-")
	query = strings.ReplaceAll(query, " ", "")

	// Log the query and the locations for debugging purposes
	//log.Printf("Searching for concerts in: %s", query)

	if id <= 0 || id > len(f.Store.Locations.Index) {
		return false
	}

	// Loop through each concert location and compare
	for _, loc := range f.Store.Locations.Index[id-1].Locations {
		// Normalize the location by trimming spaces and converting to lowercase
		normalizedLoc := strings.ToLower(strings.TrimSpace(loc))
		// Replace commas with hyphens in the location string for consistent matching
		normalizedLoc = strings.ReplaceAll(normalizedLoc, ",", "-")

		// Check if the query matches the location
		//fmt.Println(normalizedLoc, query)
		if strings.Contains(normalizedLoc, query) {
			return true
		}
	}
	return false
}
