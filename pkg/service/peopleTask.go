package service

import "github.com/TrackTasks/pkg/repository"

type PeopleTaskService struct {
	repo repository.Tasks
}

func NewPeopleTaskService(repo repository.Tasks) *PeopleTaskService {
	return &PeopleTaskService{repo: repo}
}
