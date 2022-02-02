package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"taxiTestTask/models"
)

const (
	IdCounter = "id_counter"
	globalID  = "GID_"
	ID        = "ID_"
)

func SetHashFromStruct(rdb *redis.Client, data *models.TaxiData, key string) {
	rdb.HSet(context.Background(), key, "name", data.Name, "admArea", data.AdmArea, "district", data.District,
		"address", data.Address, "carCapacity", data.CarCapacity, "mode", data.Mode, "global_id", data.GlobalId,
		"coordX", data.CoordX, "coordY", data.CoordY)
}
