// api/fetch_dates.go
package api

import (
	"encoding/json"
	"groupie-tracker/models"
	"io"
	"net/http"
)

func Fetch_dates(url string) (*models.Date, error) {
	var dates *models.Date

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &dates); err != nil {
		return nil, err
	}

	return dates, nil
}
