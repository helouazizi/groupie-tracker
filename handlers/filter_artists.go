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
	err := json.NewDecoder(r.Body).Decode(&filterReq)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}
	var result []models.Artist
	fmt.Println("Filter parameters:", filterReq)
	f.filter(filterReq, &result)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "error finging the filtered info: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func (f *Filter_Handler) filter(data FilterRequest, target *[]models.Artist) {
	//lets range over the artist repo
	from, _ := strconv.Atoi(data.CreationFrom)
	to, _ := strconv.Atoi(data.CreationTo)
	members, _ := strconv.Atoi(data.Members)
	//concerts := strings.Split(strings.ToLower(data.ConcertDate), " ")
	for _, artist := range f.Store.Artists {
		// lets make the consitons to math the filterd informatio
		craetionDate := artist.CreationDate
		album := strings.Split(artist.FirstAlbum, "-")[2]
		//includes := f.exist(artist.ID, concerts)
		if (craetionDate >= from && craetionDate <= to) || (album >= data.AlbumFrom && album <= data.AlbumTo) || (len(artist.Members) == members) {
			*target = append(*target, artist)
		}
	}
}

// func (f *Filter_Handler) exist(id int, consets []string) bool {
// 	for _,loc := range consets {
// 		if strings.Contains()
// 	}
// 	return false
// }
