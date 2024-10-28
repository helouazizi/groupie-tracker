package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var Artists interface{}

func FetchArtists() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
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

	if err := json.Unmarshal(body, &Artists); err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return nil
}
