package handlers

import (
	"groupie-tracker/repository"
	"net/http"
)

type Artist_Deatils struct {
	Store *repository.Store
}

func (s *Artist_Deatils) Artist_Deatil(w http.ResponseWriter, r *http.Request) {
	
}
