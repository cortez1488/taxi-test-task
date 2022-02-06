package service

import (
	"taxiTestTask/models"
	"taxiTestTask/pkg/repository"
)

const (
	dbExpTime = 86400
)

type taxiService struct {
	repo repository.TaxiParking
}

func newTaxiService(repo *repository.TaxiParking) *taxiService {
	return &taxiService{repo: *repo}
}

func (s *taxiService) GetById(id int) (*models.TaxiData, error) {
	return s.repo.GetById(id)
}

func (s *taxiService) GetByGlobalId(id int64) (*models.TaxiData, error) {
	return s.repo.GetByGlobalId(id)
}

func (s *taxiService) DeleteID(id int) (int64, error) {
	return s.repo.DeleteID(id)
}

func (s *taxiService) DeleteGID(id int64) (int64, error) {
	return s.repo.DeleteGID(id)
}
