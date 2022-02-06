package service

import (
	"taxiTestTask/models"
	"taxiTestTask/pkg/repository"
)

type DBLogicService struct {
	repo repository.DBLogic
}

func newDBLogicService(DBLogic repository.DBLogic) *DBLogicService {
	return &DBLogicService{repo: DBLogic}
}

func (s *DBLogicService) FillDB(data []models.TaxiData) error {
	return s.repo.FillDB(data)
}

func (s *DBLogicService) FlushDB() {
	s.repo.FlushDB()
}

func (s *DBLogicService) GetExpTimeDb() (int, error) {
	return s.repo.GetExpTimeDb()
}

func (s *DBLogicService) IncrExpTimeDb() {
	s.repo.IncrExpTimeDb()
}

func (s *DBLogicService) FreshExpTimeDb() {
	s.repo.FreshExpTimeDb()
}
