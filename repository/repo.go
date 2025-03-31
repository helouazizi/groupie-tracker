package repository

import (
	"fmt"
	"groupie-tracker/api"
	"groupie-tracker/models"
	"strconv"
	"sync"
)

// "artists": "https://groupietrackers.herokuapp.com/api/artists",
// "locations": "https://groupietrackers.herokuapp.com/api/locations",
// "dates": "https://groupietrackers.herokuapp.com/api/dates",
// "relation": "https://groupietrackers.herokuapp.com/api/relation"

type Store struct {
	Artists   []models.Artist
	Locations models.LocationsApiRespnse
	Realtions models.RelationsApiResponse
	Dates     models.DatesApiResponse
	Wg        sync.WaitGroup
	Mutex     sync.Mutex
}

func New_Store() *Store {
	return &Store{}
}

// loadd data
func (s *Store) LoadData() {
	apiUrls := []string{
		"https://groupietrackers.herokuapp.com/api/artists",
		"https://groupietrackers.herokuapp.com/api/locations",
		"https://groupietrackers.herokuapp.com/api/dates",
		"https://groupietrackers.herokuapp.com/api/relation",
	}
	s.Wg.Add(len(apiUrls))
	go api.Fetch(apiUrls[0], &s.Artists, &s.Wg)
	go api.Fetch(apiUrls[1], &s.Locations, &s.Wg)
	go api.Fetch(apiUrls[2], &s.Dates, &s.Wg)
	go api.Fetch(apiUrls[3], &s.Realtions, &s.Wg)
	s.Wg.Wait()
	fmt.Println("Fetched Artists Data:", s.Artists[0])
	//fmt.Println("Fetched Locations Data:", s.Locations.Index)
	//fmt.Println("Fetched Dates Data:", s.Dates.Index)
	//fmt.Println("Fetched Relation Data:", s.Realtions.Index)
}
func (s *Store) GetArtists() []models.Artist {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Artists
}

func (s *Store) GetArtistByID(id string) (models.Artist, bool) {
	idint, err := strconv.Atoi(id)
	if err != nil || idint < 1 || idint > len(s.Artists) {
		return models.Artist{}, false
	}
	artist := s.Artists[idint-1]
	return artist, true
}
