package service

import (
	"github.com/TrackTasks/models"
	"github.com/TrackTasks/pkg/repository"
	"github.com/sirupsen/logrus"
)

type TaskService struct {
	repo repository.Tasks
}

func NewTaskService(repo repository.Tasks) *TaskService {
	return &TaskService{repo: repo}
}
func (s *TaskService) CreateTask(newTask models.Task) (int, error) {
	if s.repo == nil {
		logrus.Error("PeopleService.Create: repository is nil")

	}
	return s.repo.CreateTask(newTask)
}
func (s *TaskService) UpdateTask(id int, task models.Task) error {
	return s.repo.UpdateTask(id, task)
}
func (s *TaskService) DeleteTask(id int, status string) (int64, error) {
	return s.repo.DeleteTask(id, status)
}
func (s *TaskService) GetTask(taskIdInt int, status string) ([]models.Task, error) {
	return s.repo.GetTask(taskIdInt, status)
}
