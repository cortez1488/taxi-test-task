package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"taxiTestTask/models"
)

type taxiParkingRedis struct {
	rdb *redis.Client
}

func newTaxiParkingRedis(rdb *redis.Client) *taxiParkingRedis {
	return &taxiParkingRedis{rdb: rdb}
}

func (r *taxiParkingRedis) Create(data *models.TaxiData) error {
	key, err := getHashCreatingKey(data, r.rdb)
	if err != nil {
		return err
	}

	_, err = r.rdb.Pipelined(context.Background(), func(rdb redis.Pipeliner) error {
		r.rdb.HSet(context.Background(), key, "name", data.Name)
		r.rdb.HSet(context.Background(), key, "admArea", data.AdmArea)
		r.rdb.HSet(context.Background(), key, "district", data.District)
		r.rdb.HSet(context.Background(), key, "address", data.Address)
		r.rdb.HSet(context.Background(), key, "carCapacity", data.CarCapacity)
		r.rdb.HSet(context.Background(), key, "mode", data.Mode)
		r.rdb.HSet(context.Background(), key, "global_id", data.GlobalId)
		r.rdb.HSet(context.Background(), key, "coordX", data.CoordX)
		r.rdb.HSet(context.Background(), key, "coordY", data.CoordY)
		return nil
	})
	if err != nil {
		return errors.New("putting in redis' hash struct data: " + err.Error())
	}
	return nil
}

func getHashCreatingKey(data *models.TaxiData, rdb *redis.Client) (string, error) {
	id, err := rdb.Incr(context.Background(), "id_counter").Result()
	if err == redis.Nil {
		return "", errors.New("can not find db's id_counter: " + err.Error())
	}
	key := fmt.Sprintf("taxi:ID_%d:GID_%d", id, data.GlobalId)
	return key, nil
}

func (r *taxiParkingRedis) Delete() {
	//sd
}
func (r *taxiParkingRedis) GetById(id int) (*models.TaxiData, error) {
	key, err := getKeyForId(id, r.rdb)
	if err != nil {
		return nil, errors.New("key for does not ex: " + err.Error())
	}
	var output models.TaxiData
	err = r.rdb.HGetAll(context.Background(), key).Scan(&output)
	if err != nil {
		return nil, errors.New("error at scanning into service struct: " + err.Error())
	}
	return &output, nil
}

func getKeyForId(id int, rdb *redis.Client) (string, error) {
	keys, err := rdb.Keys(context.Background(), fmt.Sprintf("*:ID_%d:*", id)).Result()
	if err == redis.Nil {
		return "", err
	}
	if len(keys) > 1 {
		return "", errors.New("more than 1 object with ID")
	}
	return keys[0], nil

}
func (r *taxiParkingRedis) GetByGlobalId() {
	//sd
}
