package handlers

import "groupie-tracker/repository"

type FilterRequest struct {
	CreationFrom int      `json:"creationFrom"`
	CreationTo   int      `json:"creationTo"`
	AlbumFrom    int      `json:"albumFrom"`
	AlbumTo      int      `json:"albumTo"`
	Members      int      `json:"members"`
	Locations    []string `json:"locations"`
}

type Filter_Handler struct {
	Store *repository.Store
}

func (f *Filter_Handler) Filter() {

}
