package service

import (
	"github.com/TrackTasks/models"
	"github.com/TrackTasks/pkg/repository"
)

type People interface {
	Create(peopleData models.People) (int, error)
	Update(id int, peopleData models.People) error
	GetAll() ([]models.People, error)
	GetWithFilters(id int, name string, surname string, address string, passportNumber int) ([]models.People, error)
	Delete(id int) (int64, error)
}

type Tasks interface {
	CreateTask(task models.Task) (int, error)
	GetTask(id int, status string) ([]models.Task, error)
	UpdateTask(id int, task models.Task) error
	DeleteTask(id int, status string) (int64, error)
}

type TaskTime interface {
}

type Service struct {
	People
	Tasks
	TaskTime
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		People: NewPeopleService(repos.People),
		Tasks:  NewTaskService(repos.Tasks),
	}
}
