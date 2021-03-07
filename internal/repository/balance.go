package repository

import "github.com/jmoiron/sqlx"

type Balance struct {
	db *sqlx.DB
}

func NewBalance(db *sqlx.DB) IBalance {
	return &Balance{db: db}
}
