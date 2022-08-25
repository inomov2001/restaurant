package postgres

import "database/sql"

type PostgresRepo struct {
	db *sql.DB
}

func New(db *sql.DB) PostgresRepo {
	return PostgresRepo{
		db: db,
	}
}
