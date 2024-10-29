// api/relations.go
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
	Relations *[]models.Relation
)

func init() {
	var err error
	Relations, err = fetchRelations()
	if err != nil {
		log.Fatalf("Error initializing Relations: %v", err)
	}
}

func fetchRelations() (*[]models.Relation, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
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

	var relations []models.Relation
	if err := json.Unmarshal(body, &relations); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return &relations, nil
}
