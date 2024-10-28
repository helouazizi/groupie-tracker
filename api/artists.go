// api/artists.go
package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io"
	"net/http"
	"sync"
)

var (
	Artists []models.Artist
	wg      sync.WaitGroup
)

func fetchArtistsData(url string) error {
	defer wg.Done() // Mark this goroutine as done when it finishes
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching artists: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal JSON into the Artists slice
	if err := json.Unmarshal(body, &Artists); err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return nil
}

func FetchArtists() error {
	wg.Add(1)
	go fetchArtistsData("https://groupietrackers.herokuapp.com/api/artists")

	// You can add more goroutines for other data fetching as needed
	// wg.Add(1)
	// go fetchLocationsData("...")

	wg.Wait() // Wait for all goroutines to finish
	return nil
}
