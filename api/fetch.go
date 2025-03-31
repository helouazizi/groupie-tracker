package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchJSON fetches JSON from a given URL
func Fetch(url string, target any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch: %s", url)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
