package repository

import (
	"sync"

	"go-rest-api/internal/models"
)

type Store struct {
	Artists   []models.Artist
	Locations models.LocationsApiRespnse
	Realtions models.RelationsApiResponse
	Dates     models.DatesApiResponse
	Wg        sync.WaitGroup
	Mutex     sync.Mutex
}

func New_Store() *Store {
	return &Store{}
}
