// api/fetch_details.go
package api

import (
	"groupie-tracker/models"
)

func FetchDetails(id string) (*models.Artist_Details, error) {
	//var artist_details *models.Artist_Details
	artist, err := GetArtistByID(id)
	if err != nil {
		return nil, err

	}

	locations, err := Fetch_locations(artist.Locations)
	if err != nil {

		return nil, err
	}
	/*
		relations, err := Fetch_Relations(artist.Relations)
		if err != nil {

			return nil, err
		}*/

	dates, err := Fetch_Dates(artist.ConcertDate)
	if err != nil {

		return nil, err
	}

	return &models.Artist_Details{
		Name:         artist.Name,
		Image:        artist.Image,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Members:      artist.Members,
		Locations:    locations.Locations,
		ConcertDate:  dates.Dates,
	}, nil

}
