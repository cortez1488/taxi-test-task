package app

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"taxiTestTask/internal/db_update"
	"taxiTestTask/models"
	"taxiTestTask/pkg/repository"
)

func Run() {
	db := InitRedis()
	repo := repository.NewRepositoryRedis(db)
	fmt.Sprintf("", repo)

}

func InitRedis() *redis.Client {
	testData := models.TaxiData{
		Name:        "testName2",
		AdmArea:     "testArea2",
		District:    "testDistrict2",
		Address:     "testDistrict2",
		CarCapacity: 7,
		Mode:        "круглосуточное",
		GlobalId:    40005454,
		CoordX:      13.68798,
		CoordY:      56.21545,
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	rdb.SetNX(context.Background(), repository.IdCounter, "0", 0)
	repo := repository.NewRepositoryRedis(rdb)

	err := repo.Create(&testData)
	if err != nil {
		log.Fatal(err)
	}

	result, err := repo.GetByGlobalId(40005454)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*result)
	return rdb
}

func FillRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	repo := repository.NewRepositoryRedis(rdb)
	err := db_update.RefillDB(repo)
	if err != nil {
		log.Fatal(err)
	}
}
