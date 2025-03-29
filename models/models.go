package models

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	// Members      []string `json:"members"`
	// FirstAlbum   string   `json:"firstAlbum"`
	// CreationDate int      `json:"creationDate"`
}

// type Location struct {
// 	ID        int      `json:"id"`
// 	Locations []string `json:"locations"`
// }

// type Date struct {
// 	ID    int      `json:"id"`
// 	Dates []string `json:"dates"`
// }

// type Relation struct {
// 	ID       int                 `json:"id"`
// 	Relation map[string][]string `json:"relation"`
// }
