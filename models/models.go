// models/models.go
// models.go
package models

type Artist struct {
	// Define the fields based on the API response
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	Image        string   `json:"image"`       // api
	Locations    string   `json:"locations"`   // api
	Relations    string   `json:"relations"`   // api
	ConcertDate  string   `json:"concertDate"` // api
}
type Artist_Details struct {
	// Define the fields based on the API response
	ID           int
	Name         string
	CreationDate int
	FirstAlbum   string
	Image        string
	Members      []string
	Locations    []string
	Relations    []string
	ConcertDate  []string
	Dates        []string
}
type Location struct {
	// Define fields
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
