package service

import "github.com/TrackTasks/pkg/repository"

type People interface {
}

type Tasks interface {
}

type TaskTime interface {
}

type Service struct {
	People
	Tasks
	TaskTime
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
