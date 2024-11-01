package repository

import "github.com/jmoiron/sqlx"

type PeopleTaskPostgres struct {
	db *sqlx.DB
}

func NewPeopleTaskPostgres(db *sqlx.DB) *PeopleTaskPostgres {
	return &PeopleTaskPostgres{db: db}
}
