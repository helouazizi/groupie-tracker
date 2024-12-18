// api/fetch_dates.go
// api/fetch_relations.go
package api

import (
	"encoding/json"
	"groupie-tracker/models"
	"io"
	"net/http"
)

func Fetch_Dates(url string) (*models.Date, error) {
	dates := &models.Date{}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, dates); err != nil {
		return nil, err
	}

	return dates, nil

}
