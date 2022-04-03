package repository

import (
	"clother"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ItemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) CreateItem(item clother.Item) (uuid.UUID, error) {
	var id uuid.UUID

	query := "INSERT INTO items (title, description, cost, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	row := r.db.QueryRowx(query, item.Title, item.Description, item.Cost, item.CategoryID)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *ItemRepository) GetAllCategoryItems(id uuid.UUID) ([]clother.Item, error) {
	var items []clother.Item

	err := r.db.Select(&items, "SELECT * FROM items WHERE category_id = $1", id)

	return items, err
}
