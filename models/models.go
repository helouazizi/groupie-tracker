// models/models.go
// models.go
package models

/*
type ApiResponse struct {
	Artists   string `json:"artists"`   // pai
	Locations string `json:"locations"` // api
	Dates     string `json:"dates"`     // api
	Relation  string `json:"relation"`  // api
}*/
type Artist struct {
	// Define the fields based on the API response
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	Image        string   `json:"image"` // api
	Locations    string   `json:"locations"`
	Local        struct{ Location }
	Relations    string `json:"relations"`
	Rel          struct{ Relation }
	ConcertDate  string `json:"concertDates"`
	Dat          struct{ Date }
}

type Artist_Details struct {
	// Define the fields based on the API response
	ID             int
	Name           string
	CreationDate   int
	FirstAlbum     string
	Image          string
	Members        []string
	Locations      []string
	ConcertDate    []string
	DatesLocations map[string][]string
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
