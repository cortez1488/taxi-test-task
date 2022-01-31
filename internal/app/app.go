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
	db := InitRedis()
	repo := repository.NewRepositoryRedis(db)
	fmt.Sprintf("", repo)

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

func InitRedis() *redis.Client {
	//testData := models.TaxiData{
	//	Name:        "testName2",
	//	AdmArea:     "testArea2",
	//	District:    "testDistrict2",
	//	Address:     "testDistrict2",
	//	CarCapacity: 7,
	//	Mode:        "круглосуточное",
	//	GlobalId:    1488661122,
	//	CoordX:      13.68798,
	//	CoordY:      56.21545,
	//}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	rdb.SetNX(context.Background(), "id_counter", "0", 0)
	repo := repository.NewRepositoryRedis(rdb)

	//err := repo.Create(&testData)
	//if err != nil {
	//	log.Fatal(err)
	//}

	result, err := repo.DeleteGID(888888888)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	return rdb
}
