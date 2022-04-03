package repository

import (
	"clother"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r CategoryRepository) CreateCategory(category clother.Category) (uuid.UUID, error) {
	var id uuid.UUID

	query := "INSERT INTO categories (title) VALUES ($1) RETURNING id"
	row := r.db.QueryRowx(query, category.Title)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *CategoryRepository) GetAllCategories() ([]clother.Category, error) {
	var categories []clother.Category
	err := r.db.Select(&categories, "SELECT * FROM categories")

	return categories, err
}

func (r *CategoryRepository) GetCategoryByID(id uuid.UUID) (clother.Category, error) {
	var category clother.Category
	err := r.db.Get(&category, "SELECT * FROM categories WHERE id = $1", id)

	return category, err
}
