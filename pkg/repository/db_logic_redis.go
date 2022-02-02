package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"taxiTestTask/models"
)

type DBLogicRedis struct {
	rdb *redis.Client
}

func NewDBLogicRedis(rdb *redis.Client) *DBLogicRedis {
	return &DBLogicRedis{rdb: rdb}
}

func (r *DBLogicRedis) FlushDB() {
	r.rdb.FlushDB(context.Background())
}

func (r *DBLogicRedis) FillDB(slice *[]models.TaxiData) error {
	r.rdb.SetNX(context.Background(), IdCounter, "0", 0)
	for _, data := range *slice {
		key, err := getHashCreatingKey(&data, r.rdb)
		if err != nil {
			return err
		}
		_, err = r.rdb.Pipelined(context.Background(), func(rdb redis.Pipeliner) error {
			SetHashFromStruct(r.rdb, &data, key)
			return nil
		})
	}
	return nil
}

func (r *DBLogicRedis) GetExpTimeDb() (int, error) {
	return r.rdb.Get(context.Background(), timeRefillDB).Int()
}

func (r *DBLogicRedis) FreshExpTimeDb() {
	r.rdb.Set(context.Background(), timeRefillDB, 0, 0)
}

func (r *DBLogicRedis) IncrExpTimeDb() {
	r.rdb.Incr(context.Background(), timeRefillDB)
}
