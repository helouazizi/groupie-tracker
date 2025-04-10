package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/repository"
	"net/http"
)

type FilterRequest struct {
	CreationFrom int `json:"creationDateFrom"`
	CreationTo   int `json:"creationDateTo"`
	AlbumFrom    int `json:"firstAlbumFrom"`
	AlbumTo      int `json:"firstAlbumTo"`
	Members      int `json:"members"`
	//Locations    []string `json:"locations"`
}

type Filter_Handler struct {
	Store *repository.Store
}

func (f *Filter_Handler) Filter(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for preflight and actual request
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// 🧠 Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var filterReq FilterRequest
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&filterReq)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Filter parameters:", filterReq)

	// Here you would apply your filtering logic using f.Store
	// filtered := f.Store.FilterArtists(filterReq)
	// json.NewEncoder(w).Encode(filtered)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Filter received!"})
}
