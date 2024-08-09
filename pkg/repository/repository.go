package repository

import "github.com/jmoiron/sqlx"

type People interface {
}

type Tasks interface {
}

type TaskTime interface {
}

type Repository struct {
	People
	Tasks
	TaskTime
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
