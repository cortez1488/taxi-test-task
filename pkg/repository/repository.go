package repository

import (
	"github.com/go-redis/redis/v8"
	"taxiTestTask/models"
)

type TaxiParking interface {
	Create(data *models.TaxiData) error
	GetById(id int) (*models.TaxiData, error)
}

type Repository struct {
	TaxiParking
}

func NewRepositoryRedis(rdb *redis.Client) *Repository {
	return &Repository{TaxiParking: newTaxiParkingRedis(rdb)}
}
