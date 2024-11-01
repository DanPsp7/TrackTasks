package repository

import (
	"github.com/TrackTasks/models"
	"github.com/jmoiron/sqlx"
)

type People interface {
	Create(newPeople models.People) (int, error)
	Update(id int, updatedPeople models.People) error
	GetAll() ([]models.People, error)
	GetWithFilters(id int, name string, surname string, address string, passportNumber int) ([]models.People, error)
	Delete(id int) (int64, error)
}

type Tasks interface {
	CreateTask(newTask models.Task) (int, error)
	UpdateTask(id int, updatedTask models.Task) error
	GetTask(id int, status string) ([]models.Task, error)
	DeleteTask(id int, status string) (int64, error)
}

type TaskTime interface {
	StartTask(id int) (models.Task, error)
	StopTask(id int) (models.Task, error)
}

type Repository struct {
	People
	Tasks
	TaskTime
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		People: NewPeoplePostgres(db),
		Tasks:  NewTaskPostgres(db),
	}
}
