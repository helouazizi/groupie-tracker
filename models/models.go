package models

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	Image        string   `json:"image"`     // api
	Locations    string   `json:"locations"` // api
	Local        struct{ Location }
	Relations    string `json:"relations"` // api
	Rel          struct{ Relation }
	ConcertDate  string `json:"concertDates"` // api
	Dat          struct{ Date }
}

type Artist_Details struct {
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
