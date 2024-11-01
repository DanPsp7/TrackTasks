package service

import (
	"github.com/TrackTasks/models"
	"github.com/TrackTasks/pkg/repository"
	"github.com/sirupsen/logrus"
)

type PeopleService struct {
	repo repository.People
}

func NewPeopleService(repo repository.People) *PeopleService {
	return &PeopleService{repo: repo}
}
func (s *PeopleService) Create(newPeople models.People) (int, error) {
	if s.repo == nil {
		logrus.Error("PeopleService.Create: repository is nil")

	}
	return s.repo.Create(newPeople)
}

func (s *PeopleService) Update(id int, updatedPeople models.People) error {
	return s.repo.Update(id, updatedPeople)

}

func (s *PeopleService) GetAll() ([]models.People, error) {
	return s.repo.GetAll()
}
func (s *PeopleService) GetWithFilters(id int, name string, surname string, address string, passportNumber int) ([]models.People, error) {
	return s.repo.GetWithFilters(id, name, surname, address, passportNumber)
}

func (s *PeopleService) Delete(id int) (int64, error) {
	return s.repo.Delete(id)
}
