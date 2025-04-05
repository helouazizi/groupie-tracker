package handlers

import (
	"encoding/json"
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

}
