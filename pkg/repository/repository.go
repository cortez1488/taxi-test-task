package repository

import "taxiTestTask/models"

type TaxiParking interface {
	Create(data *models.TaxiData) error
}

type Repository struct {
	TaxiParking
}
