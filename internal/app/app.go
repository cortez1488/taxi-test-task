package app

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"taxiTestTask/internal/json_to_struct"
	reqToAPI "taxiTestTask/internal/reqToAPI/JSON"
	"taxiTestTask/models"
	"taxiTestTask/pkg/repository"
)

func Run() {

	json, err := reqToAPI.GetJSONFromAPIRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	var input []models.TaxiData
	err = json_to_struct.Parse(json, &input)
	if err != nil {
		log.Fatal(err.Error())
	}
	for i, taxi := range input {
		fmt.Println(i, taxi.Mode)
	}
}

func Redis() {
	testData := models.TaxiData{
		Name:        "testName",
		AdmArea:     "testArea",
		District:    "testDistrict",
		Address:     "testAddress",
		CarCapacity: 5,
		Mode:        "круглосуточное",
		GlobalId:    888888888,
		Coords:      []float32{54.1256, 65.6546454},
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	repo := repository.NewRepositoryRedis(rdb)
	rdb.SetNX(context.Background(), "id_counter", "0", 0)

	err := repo.TaxiParking.Create(&testData)
	if err != nil {
		log.Fatal(err.Error())
	}
}
