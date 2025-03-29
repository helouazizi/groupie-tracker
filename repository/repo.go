package repository

import (
	"groupie-tracker/models"
	"sync"
)

type Store struct {
	Artists []models.Artist
	Mutex   sync.Mutex
}
