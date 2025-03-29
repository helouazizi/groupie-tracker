package repository

import (
	"groupie-tracker/models"
	"sync"
)

type Store struct {
	Artists []models.Artist
	Mutex   sync.Mutex
}

func New_Store() *Store {
	return &Store{}
}


