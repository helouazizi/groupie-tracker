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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Fix CORS issues
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var filterReq FilterRequest
	err := json.NewDecoder(r.Body).Decode(&filterReq)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// You now have the data
	fmt.Println("Filter parameters:", filterReq)

	// TODO: use the data to filter artists and return them as JSON
	// e.g., filtered := filterArtists(filterReq)
	// json.NewEncoder(w).Encode(filtered)

	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Filter received!"})

}
