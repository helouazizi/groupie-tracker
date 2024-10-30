// api/fetch_details.go
package api

import (
	"fmt"
	"log"
)

func FetchDetails(id string) ([]string, error) {
	artist, err := GetArtistByID(id)
	if err != nil {
		log.Fatal("cant find artist")
	}
	fmt.Println(artist.Locations)
	locations, err := Fetch_locations(artist.Locations)
	if err != nil {
		log.Fatal("cant find locations")
	}

	return locations, nil

}
