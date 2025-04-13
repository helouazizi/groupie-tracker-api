package services

import (
	"go-rest-api/internal/models"
	"go-rest-api/internal/repository"
)

type FilterService struct {
	Store *repository.Store
}

func NewFilterService(store *repository.Store) *FilterService {
	return &FilterService{Store: store}
}

func (s *FilterService) Filter(data models.FilterRequest) []models.Artist {
	var filteredartists []models.Artist
	return filteredartists
}
