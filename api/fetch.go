package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
)

// FetchJSON fetches JSON from a given URL
func Fetch(url string, target interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching data from %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch data from %s: StatusCode %d", url, resp.StatusCode)
		return
	}
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body from %s: %v", url, err)
		return
	}
	// Parse the JSON response into the appropriate struct
	if err := json.Unmarshal(body, &target); err != nil {
		log.Printf("Error unmarshalling data from %s: %v", url, err)
		return
	}
}
