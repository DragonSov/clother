package service

import (
	"clother"
	"clother/pkg/repository"
	"github.com/google/uuid"
)

type Authorization interface {
	CreateUser(clother.User) (uuid.UUID, error)
	GetUserByID(uuid.UUID) (clother.User, error)
	GetUserByLogin(string) (clother.User, error)
	GenerateToken(string, string) (string, error)
	ParseToken(string) (uuid.UUID, error)
}

type Category interface {
	CreateCategory(clother.Category) (uuid.UUID, error)
	GetAllCategories() ([]clother.Category, error)
	GetCategoryByID(uuid.UUID) (clother.Category, error)
}

type Item interface {
	CreateItem(clother.Item) (uuid.UUID, error)
	GetAllCategoryItems(uuid.UUID) ([]clother.Item, error)
}

type Service struct {
	Authorization
	Category
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Category:      NewCategoryService(repos),
		Item:          NewItemService(repos),
	}
}
