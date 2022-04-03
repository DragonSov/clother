package clother

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Login     string    `db:"login"`
	Password  string    `db:"password"`
	Admin     bool      `db:"admin"`
	CreatedAt time.Time `db:"created_at"`
}

type Category struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}

type Item struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Cost        int       `db:"cost"`
	CategoryID  uuid.UUID `db:"category_id"`
	CreatedAt   time.Time `db:"created_at"`
}
