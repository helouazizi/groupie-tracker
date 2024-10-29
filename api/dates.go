// api/dates.go
// api/dates.go
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
	Dates *[]models.Date
)

func init() {
	var err error
	Dates, err = fetchDates()
	if err != nil {
		log.Fatalf("Error initializing Dates: %v", err)
	}
}

func fetchDates() (*[]models.Date, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
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

	var dates []models.Date
	if err := json.Unmarshal(body, &dates); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return &dates, nil
}