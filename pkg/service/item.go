package service

import (
	"clother"
	"clother/pkg/repository"
	"github.com/google/uuid"
)

type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem(item clother.Item) (uuid.UUID, error) {
	return s.repo.CreateItem(item)
}

func (s *ItemService) GetAllCategoryItems(id uuid.UUID) ([]clother.Item, error) {
	return s.repo.GetAllCategoryItems(id)
}
