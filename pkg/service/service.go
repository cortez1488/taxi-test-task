package service

import (
	"taxiTestTask/models"
	"taxiTestTask/pkg/repository"
)

type TaxiParking interface {
	GetById(id int) (*models.TaxiData, error)
	GetByGlobalId(globalId int64) (*models.TaxiData, error)
	DeleteID(id int) (int64, error)
	DeleteGID(id int64) (int64, error)
	RefillDB(*[]models.TaxiData) error
}

type Service struct {
	TaxiParking
}

func NewService(repo *repository.Repository) *Service {
	return &Service{TaxiParking: newTaxiService(&repo.TaxiParking)}
}
