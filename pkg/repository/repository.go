package repository

import (
	"clother"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(clother.User) (uuid.UUID, error)
	GetUserByID(uuid.UUID) (clother.User, error)
	GetUserByLogin(string) (clother.User, error)
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

type Repository struct {
	Authorization
	Category
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Category:      NewCategoryRepository(db),
		Item:          NewItemRepository(db),
	}
}
