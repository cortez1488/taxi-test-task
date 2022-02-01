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
		r.setHashFromStruct(data, key)
		return nil
	})
	if err != nil {
		return errors.New("putting in redis' hash struct data: " + err.Error())
	}
	return nil
}

func (r *taxiParkingRedis) FlushDB() {
	r.rdb.FlushDB(context.Background())
}

func (r *taxiParkingRedis) FillDB(slice *[]models.TaxiData) error {
	r.rdb.SetNX(context.Background(), IdCounter, "0", 0)
	for _, data := range *slice {
		key, err := getHashCreatingKey(&data, r.rdb)
		if err != nil {
			return err
		}
		_, err = r.rdb.Pipelined(context.Background(), func(rdb redis.Pipeliner) error {
			r.setHashFromStruct(&data, key)
			return nil
		})
	}
	return nil
}

func (r *taxiParkingRedis) setHashFromStruct(data *models.TaxiData, key string) {
	r.rdb.HSet(context.Background(), key, "name", data.Name, "admArea", data.AdmArea, "district", data.District,
		"address", data.Address, "carCapacity", data.CarCapacity, "mode", data.Mode, "global_id", data.GlobalId,
		"coordX", data.CoordX, "coordY", data.CoordY)
}

func getHashCreatingKey(data *models.TaxiData, rdb *redis.Client) (string, error) {
	id, err := rdb.Incr(context.Background(), IdCounter).Result()
	if err == redis.Nil {
		return "", errors.New("can not find db's id counter: " + err.Error())
	}
	key := fmt.Sprintf("taxi:%s%d:%s%d:", ID, id, globalID, data.GlobalId)
	return key, nil
}

func (r *taxiParkingRedis) GetById(id int) (*models.TaxiData, error) {
	key, err := getKeyForId(id, r.rdb)
	if err != nil {
		return nil, errors.New("key id does not ex: " + err.Error())
	}
	var output models.TaxiData
	err = r.rdb.HGetAll(context.Background(), key).Scan(&output)
	if err != nil {
		return nil, errors.New("error at scanning into service struct: " + err.Error())
	}
	return &output, nil
}

func (r *taxiParkingRedis) GetByGlobalId(globalId int64) (*models.TaxiData, error) {
	key, err := getKeyForGlobalId(globalId, r.rdb)
	if err != nil {
		return nil, errors.New("key gid error: " + err.Error())
	}
	var output models.TaxiData
	err = r.rdb.HGetAll(context.Background(), key).Scan(&output)
	if err != nil {
		return nil, errors.New("error at scanning into service struct: " + err.Error())
	}
	return &output, nil
}

func (r *taxiParkingRedis) DeleteID(id int) (int64, error) {
	key, err := getKeyForId(id, r.rdb)
	if err != nil {
		return 0, err
	}
	return r.rdb.Del(context.Background(), key).Val(), nil
}

func (r *taxiParkingRedis) DeleteGID(id int64) (int64, error) {
	key, err := getKeyForGlobalId(id, r.rdb)
	if err != nil {
		return 0, err
	}
	return r.rdb.Del(context.Background(), key).Val(), nil
}

func getKeyForId(id int, rdb *redis.Client) (string, error) {
	keys, err := rdb.Keys(context.Background(), fmt.Sprintf("*:%s%d:*", ID, id)).Result()
	if err == redis.Nil {
		return "", err
	}
	if len(keys) > 1 {
		return "", errors.New("more than 1 object with ID")
	}
	return keys[0], nil

}

func getKeyForGlobalId(id int64, rdb *redis.Client) (string, error) {
	keys, err := rdb.Keys(context.Background(), fmt.Sprintf("*:%s%d*", globalID, id)).Result()
	if len(keys) < 1 {
		return "", errors.New("key doesn't exists" + err.Error())
	}
	if len(keys) > 1 {
		return "", errors.New("more than 1 object with GID")
	}
	return keys[0], nil

}
