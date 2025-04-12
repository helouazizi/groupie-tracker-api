package services

import (
	"sync"

	"go-rest-api/internal/models"
)

type ItemService struct {
	Mu     sync.Mutex
	Items  map[int]models.Item
	NextId int
}

func NewItemService() *ItemService {
	return &ItemService{Items: make(map[int]models.Item), NextId: 1}
}

func (s *ItemService) GetAllItems() []models.Item {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	result := []models.Item{}
	for _, item := range s.Items {
		result = append(result, item)
	}
	return result
}

func (s *ItemService) GetItemById(id int) (models.Item, bool) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	item, exist := s.Items[id]
	return item, exist
}

func (s *ItemService) CraeteItem(name string) models.Item {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	item := models.Item{
		ID:   s.NextId,
		Name: name,
	}
	s.Items[s.NextId] = item
	s.NextId++
	return item
}

func (s *ItemService) DeleteItem(id int) bool {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	if _, exists := s.Items[id]; exists {
		delete(s.Items, id)
		return true
	}
	return false
}
