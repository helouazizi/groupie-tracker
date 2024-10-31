// api/fetch_locations.go
package api

import (
	"encoding/json"
	"groupie-tracker/models"
	"io"
	"net/http"
)

func Fetch_locations(url string) (*models.Location, error) {

	locations := &models.Location{}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, locations); err != nil {
		return nil, err
	}

	return locations, nil

}
