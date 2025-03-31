package repository

import (
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
	Locations []models.Location
	Realtions []models.Relation
	Dates     []models.Date
	Wg        sync.WaitGroup
	Mutex     sync.Mutex
}

func New_Store() *Store {
	return &Store{}
}

// loadd data
func (s *Store) LoadData() {
	// // lets request the api to get artist data
	// url := "https://groupietrackers.herokuapp.com/api/artists"
	// res, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()

	// content, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Unmarshal JSON into the struct
	// if err = json.Unmarshal(content, &s.Artists); err != nil {
	// 	log.Fatal(err)
	// }
	////////////////

	// for i, artist := range s.Artists {
	// 	//url1 := "https://groupietrackers.herokuapp.com/api/locations"
	// 	res, errr := http.Get(artist.Locations)
	// 	if errr != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer res.Body.Close()

	// 	contentt, err := io.ReadAll(res.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	var locres LocationsResponse
	// 	// Unmarshal JSON into the struct
	// 	if err = json.Unmarshal(contentt, &locres); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	s.Artists[i].LocationArray = locres.Index
	// }

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
