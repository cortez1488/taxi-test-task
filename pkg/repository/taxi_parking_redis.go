package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"taxiTestTask/models"
	"taxiTestTask/pkg/service_errors"
)

const (
	timeRefillDB = "timeRefillDB"
)

type taxiParkingRedis struct {
	rdb *redis.Client
}

func newTaxiParkingRedis(rdb *redis.Client) *taxiParkingRedis {
	return &taxiParkingRedis{rdb: rdb}
}

func (r *taxiParkingRedis) GetById(id int) (models.TaxiData, error) {
	var output models.TaxiData
	key, err := getKeyForId(id, r.rdb)
	if err != nil {
		return output, err
	}
	if key == "" {
		return output, service_errors.ErrNoData
	}

	err = r.rdb.HGetAll(context.Background(), key).Scan(&output)
	if err != nil {
		return output, errors.New("error at scanning into service struct: " + err.Error())
	}
	return output, nil
}

func (r *taxiParkingRedis) GetByGlobalId(globalId int64) (models.TaxiData, error) {
	var output models.TaxiData

	key, err := getKeyForGlobalId(globalId, r.rdb)
	if err != nil {
		return output, errors.New("key gid error: " + err.Error())
	}
	if key == "" {
		return output, service_errors.ErrNoData
	}

	err = r.rdb.HGetAll(context.Background(), key).Scan(&output)
	if err != nil {
		return output, errors.New("error at scanning into service struct: " + err.Error())
	}
	return output, nil
}

func (r *taxiParkingRedis) DeleteID(id int) (int64, error) {
	key, err := getKeyForId(id, r.rdb)
	if err != nil {
		return 0, err
	}
	if key == "" { //Error is nil, but string is empty(cause no keys)
		return 0, nil
	}
	return r.rdb.Del(context.Background(), key).Val(), nil
}

func (r *taxiParkingRedis) DeleteGID(id int64) (int64, error) {
	key, err := getKeyForGlobalId(id, r.rdb)
	if err != nil {
		return 0, err
	}
	if key == "" { //Error is nil, but string is empty(cause no keys)
		return 0, nil
	}
	return r.rdb.Del(context.Background(), key).Val(), nil
}

func getHashCreatingKey(data models.TaxiData, rdb *redis.Client) (string, error) {
	id, err := rdb.Incr(context.Background(), IdCounter).Result()
	if err == redis.Nil {
		return "", errors.New("can not find db's id counter: " + err.Error())
	}
	key := fmt.Sprintf("taxi:%s%d:%s%d:", ID, id, globalID, data.GlobalId)
	return key, nil
}

func getKeyForId(id int, rdb *redis.Client) (string, error) {
	keys, err := rdb.Keys(context.Background(), fmt.Sprintf("*:%s%d:*", ID, id)).Result()
	if err != nil {
		return "", err
	}
	if len(keys) < 1 {
		return "", nil
	}
	if len(keys) > 1 {
		return "", errors.New("more than 1 object with ID")
	}
	return keys[0], nil

}

func getKeyForGlobalId(id int64, rdb *redis.Client) (string, error) {
	keys, err := rdb.Keys(context.Background(), fmt.Sprintf("*:%s%d*", globalID, id)).Result()
	if err != nil {
		return "", err
	}
	if len(keys) < 1 {
		return "", nil
	}
	if len(keys) > 1 {
		return "", errors.New("more than 1 object with GID")
	}
	return keys[0], nil

}
