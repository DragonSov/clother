package repository

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    login text,
    password text,
    admin bool DEFAULT false,
    created_at timestamp DEFAULT now()
);

CREATE TABLE IF NOT EXISTS categories (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    title text,
    created_at timestamp DEFAULT now()
);

CREATE TABLE IF NOT EXISTS items (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    title text,
    description text,
    cost int,
    category_id uuid REFERENCES categories(id) ON DELETE CASCADE,
    created_at timestamp DEFAULT now()
);
`

func NewConnection(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
