package repository

import (
	"clother"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user clother.User) (uuid.UUID, error) {
	var id uuid.UUID

	query := "INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id"
	row := r.db.QueryRowx(query, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *AuthRepository) GetUserByID(id uuid.UUID) (clother.User, error) {
	var user clother.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)

	return user, err
}

func (r *AuthRepository) GetUserByLogin(login string) (clother.User, error) {
	var user clother.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE lower(login) = lower($1)", login)

	return user, err
}
