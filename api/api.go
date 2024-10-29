// api/api.go
// api/api.go
package api

import (
	"log"
)

func InitializeAllData() {
	// This function can be called to ensure all data is initialized
	if Artists == nil {
		log.Println("Artists data not initialized")
	}
	if Locations == nil {
		log.Println("Locations data not initialized")
	}
	if Dates == nil {
		log.Println("Dates data not initialized")
	}
	if Relations == nil {
		log.Println("Relations data not initialized")
	}
}
