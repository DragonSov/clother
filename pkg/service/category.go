package service

import (
	"clother"
	"clother/pkg/repository"
	"github.com/google/uuid"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category clother.Category) (uuid.UUID, error) {
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) GetAllCategories() ([]clother.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *CategoryService) GetCategoryByID(id uuid.UUID) (clother.Category, error) {
	return s.repo.GetCategoryByID(id)
}
