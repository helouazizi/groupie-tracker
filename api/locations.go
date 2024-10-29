// api/locations.go
// api/locations.go
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
	Locations *[]models.Location
)

func init() {
	var err error
	Locations, err = fetchLocations()
	if err != nil {
		log.Fatalf("Error initializing Locations: %v", err)
	}
}

func fetchLocations() (*[]models.Location, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
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

	var locations []models.Location
	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return &locations, nil
}
