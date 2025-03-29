package repository

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// "artists": "https://groupietrackers.herokuapp.com/api/artists",
// "locations": "https://groupietrackers.herokuapp.com/api/locations",
// "dates": "https://groupietrackers.herokuapp.com/api/dates",
// "relation": "https://groupietrackers.herokuapp.com/api/relation"

type Store struct {
	Artists []models.Artist
	Mutex   sync.Mutex
}

func New_Store() *Store {
	return &Store{}
}

// loadd data
func (s *Store) LoadData() {
	// lets request the api to get artist data
	url := "https://groupietrackers.herokuapp.com/api/artists"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal JSON into the struct
	if err = json.Unmarshal(content, &s.Artists); err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Artists[0].Name)
}

func (s *Store) GetArtists() []models.Artist {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Artists
}

func (s *Store) GetArtistByID(id string) (models.Artist, bool) {
	idint, err := strconv.Atoi(id)
	if err != nil {
		return models.Artist{}, false
	}
	artist := s.Artists[idint-1]
	return artist, true
}
