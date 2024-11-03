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
	Artists []models.Artist
)

func init() {

	err := fetchArtists()
	if err != nil {
		log.Fatal(err)
	}
}

func fetchArtists() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return err

	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &Artists); err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return nil

}
