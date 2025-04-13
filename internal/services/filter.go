package services

import (
	"fmt"
	"strconv"
	"strings"

	"go-rest-api/internal/models"
	"go-rest-api/internal/repository"
)

type FilterService struct {
	Store *repository.Store
}

func NewFilterService(store *repository.Store) *FilterService {
	return &FilterService{Store: store}
}

func (s *FilterService) Filter(data models.FilterRequest) ([]models.Artist, error) {
	var filteredartists []models.Artist
	// Parse filter values
	creationFrom, err1 := strconv.Atoi(data.CreationFrom)
	creationTo, err2 := strconv.Atoi(data.CreationTo)
	albumFrom, err3 := strconv.Atoi(data.AlbumFrom)
	albumTo, err4 := strconv.Atoi(data.AlbumTo)
	members, err5 := strconv.Atoi(data.Members)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		return nil, fmt.Errorf("error parsing number atoi error")
	}

	// Normalize the concert query
	concerts := strings.ToLower(data.ConcertDate)

	// lets range over our artists
	for _, artist := range s.Store.Artists {
		///////////////////////////
		if creationFrom != 0 && artist.CreationDate < creationFrom {
			continue
		}
		if creationTo != 0 && artist.CreationDate > creationTo {
			continue
		}
		//////////////////////////
		parts := strings.Split(artist.FirstAlbum, "-")
		// if len(parts) != 3 {
		// 	continue
		// }
		albumYear, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}
		if albumFrom != 0 && albumYear < albumFrom {
			continue
		}
		if albumTo != 0 && albumYear > albumTo {
			continue
		}
		///////////////////////////
		if members != 0 && (len(artist.Members) != members) {
			continue
		}
		///////////////////////////
		if concerts != "" && !s.matchLocation(artist.ID, concerts) {
			continue
		}
		filteredartists = append(filteredartists, artist)
	}
	return filteredartists, nil
}

func (f *FilterService) matchLocation(id int, query string) bool {
	query = strings.ToLower(strings.TrimSpace(query))
	query = strings.ReplaceAll(query, ",", "-")
	query = strings.ReplaceAll(query, " ", "")

	if id <= 0 || id > len(f.Store.Locations.Index) {
		return false
	}

	// Loop through each concert location and compare
	for _, loc := range f.Store.Locations.Index[id-1].Locations {
		// Normalize the location by trimming spaces and converting to lowercase
		normalizedLoc := strings.ToLower(strings.TrimSpace(loc))
		// Replace commas with hyphens in the location string for consistent matching
		normalizedLoc = strings.ReplaceAll(normalizedLoc, ",", "-")

		// Check if the query matches the location
		// fmt.Println(normalizedLoc, query)
		if strings.Contains(normalizedLoc, query) {
			return true
		}
	}
	return false
}
