// api/fetch_artists.go
// api/artists.go
package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io"
	"log"
	"net/http"
)

var (
	Artists *[]models.Artist
)

func init() {
	var err error
	Artists, err = fetchArtists()
	if err != nil {
		log.Fatalf("Error initializing Artists: %v", err)
	}
}

func fetchArtists() (*[]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}

	var artists []models.Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return &artists, nil

}
