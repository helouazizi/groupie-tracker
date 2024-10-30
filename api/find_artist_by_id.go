// api/find_artist_by_id.go
package api

import (
	"fmt"
	"groupie-tracker/models"
	"strconv"
)

func GetArtistByID(id string) (*models.Artist, error) {
	ids, _ := strconv.Atoi(id)
	for _, artist := range *Artists { // 'artists' is a slice of Artist structs
		if artist.ID == ids {
			return &artist, nil
		}
	}
	return nil, fmt.Errorf("artist not found")
}
