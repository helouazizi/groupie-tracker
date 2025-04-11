package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"groupie-tracker/repository"
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
	creationFrom, _ := strconv.Atoi(data.CreationFrom)
	creationTo, _ := strconv.Atoi(data.CreationTo)

	albumFrom, _ := strconv.Atoi(data.AlbumFrom)
	albumTo, _ := strconv.Atoi(data.AlbumTo)

	members, _ := strconv.Atoi(data.Members)

	concertQuery := strings.ToLower(data.ConcertDate)

	for _, artist := range f.Store.Artists {
		// ✅ Creation date filter
		if creationFrom != 0 && artist.CreationDate < creationFrom {
			continue
		}
		if creationTo != 0 && artist.CreationDate > creationTo {
			continue
		}

		// ✅ First album year filter
		parts := strings.Split(artist.FirstAlbum, "-")
		if len(parts) != 3 {
			continue
		}
		albumYear, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}
		if albumFrom != 0 && albumYear < albumFrom {
			continue
		}
		if albumTo != 0 && albumYear > albumTo {
			continue
		}

		// ✅ Members filter
		if members != 0 && len(artist.Members) != members {
			continue
		}

		// ✅ Concert location filter
		if concertQuery != "" && !f.matchLocation(artist.ID, concertQuery) {
			continue
		}

		// ✅ All conditions passed
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
		fmt.Println(normalizedLoc, query)
		if strings.Contains(normalizedLoc, query) {
			return true
		}
	}
	return false
}
