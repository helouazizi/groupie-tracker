// api/fetch_relations.go
package api

import (
	"encoding/json"
	"groupie-tracker/models"
	"io"
	"net/http"
)

func Fetch_Relations(url string) (*models.Relation, error) {
	relations := &models.Relation{}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, relations); err != nil {
		return nil, err
	}

	return relations, nil

}
