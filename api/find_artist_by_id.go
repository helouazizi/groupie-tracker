// api/find_artist_by_id.go
package api

import (
	"fmt"
	"strconv"

	"groupie-tracker/models"
)

func GetArtistByID(id string) (*models.Artist, error) {
	ids, _ := strconv.Atoi(id)
	// if err != nil || (ids<0 || ids>len(Artists)) {
	// 	//error	
	// }
	// fmt.Println(Artists[ids-1])
	//fmt.Println(Artists[1])
	for _, artist := range Artists { // 'artists' is a slice of Artist structs
		if artist.ID == ids {
			return &artist, nil
		}
	}
	return nil, fmt.Errorf("artist not found")
}
