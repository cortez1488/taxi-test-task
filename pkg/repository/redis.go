package repository

import "github.com/go-redis/redis/v8"

func NewRepositoryRedis(rdb *redis.Client) *Repository {
	return &Repository{TaxiParking: newTaxiParkingRedis(rdb)}
}
