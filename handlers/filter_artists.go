package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/repository"
	"net/http"
)

type FilterRequest struct {
	CreationFrom int      `json:"creationFrom"`
	CreationTo   int      `json:"creationTo"`
	AlbumFrom    int      `json:"albumFrom"`
	AlbumTo      int      `json:"albumTo"`
	Members      int      `json:"members"`
	Locations    []string `json:"locations"`
}

type Filter_Handler struct {
	Store *repository.Store
}

func (f *Filter_Handler) Filter(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Filter received!"})

}
