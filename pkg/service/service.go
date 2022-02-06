package service

import (
	"taxiTestTask/models"
	"taxiTestTask/pkg/repository"
)

type TaxiParking interface {
	GetById(id int) (models.TaxiData, error)
	GetByGlobalId(globalId int64) (models.TaxiData, error)
	DeleteID(id int) (int64, error)
	DeleteGID(id int64) (int64, error)
}

type DBLogic interface {
	FillDB([]models.TaxiData) error
	FlushDB()
	GetExpTimeDb() (int, error)
	FreshExpTimeDb()
	IncrExpTimeDb()
}

type Service struct {
	TaxiParking
	DBLogic
}

func NewService(repo repository.Repository) *Service {
	return &Service{TaxiParking: newTaxiService(repo.TaxiParking),
		DBLogic: newDBLogicService(repo.DBLogic)}
}
