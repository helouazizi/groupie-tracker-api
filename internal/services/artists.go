package services

import (
	"go-rest-api/internal/models"
	"go-rest-api/internal/repository"
)

type AllArtistsService struct {
	Store *repository.Store
}

func NewAllArtistsService(store *repository.Store) *AllArtistsService {
	return &AllArtistsService{Store: store}
}

func (a *AllArtistsService) GetAllArtists() []models.Artist {
	a.Store.Mutex.Lock()
	defer a.Store.Mutex.Unlock()
	return a.Store.Artists
}
