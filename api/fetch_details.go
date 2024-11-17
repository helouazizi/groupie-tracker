// api/fetch_details.go
package api

import (
	"sync"

	"groupie-tracker/models"
)

// fetchDetailsConcurrently handles fetching locations, relations, and dates concurrently
func fetchDetailsConcurrently(artist *models.Artist) (locations *models.Location, relations *models.Relation, dates *models.Date, err error) {
	var wg sync.WaitGroup
	// Define error variables for each fetch
	var locErr, relErr, dateErr error
	// Fetch locations
	wg.Add(1)
	go func() {
		defer wg.Done()
		locations, locErr = Fetch_locations(artist.Locations)
	}()

	// Fetch relations
	wg.Add(1)
	go func() {
		defer wg.Done()
		relations, relErr = Fetch_Relations(artist.Relations)
	}()

	// Fetch dates
	wg.Add(1)
	go func() {
		defer wg.Done()
		dates, dateErr = Fetch_Dates(artist.ConcertDate)
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Check for any errors from goroutines
	if locErr != nil {
		return nil, nil, nil, locErr
	}
	if relErr != nil {
		return nil, nil, nil, relErr
	}
	if dateErr != nil {
		return nil, nil, nil, dateErr
	}

	return locations, relations, dates, nil
}

// FetchDetails retrieves details for an artist
func FetchDetails(id string) (*models.Artist_Details, error) {
	artist, err := GetArtistByID(id)
	if err != nil {
		return nil, err
	}

	// Fetch locations, relations, and dates concurrently
	locations, relations, dates, err := fetchDetailsConcurrently(artist)
	if err != nil {
		return nil, err
	}

	return &models.Artist_Details{
		Name:           artist.Name,
		Image:          artist.Image,
		CreationDate:   artist.CreationDate,
		FirstAlbum:     artist.FirstAlbum,
		Members:        artist.Members,
		Locations:      locations.Locations,
		ConcertDate:    dates.Dates,
		DatesLocations: relations.DatesLocations,
	}, nil
}
