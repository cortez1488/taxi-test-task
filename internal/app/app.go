package app

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"taxiTestTask/internal/json_to_struct"
	"taxiTestTask/models"
	"taxiTestTask/pkg/handler"
	"taxiTestTask/pkg/repository"
	service2 "taxiTestTask/pkg/service"
	"time"
)

const (
	serverUrl = "localhost:8000"
	dbExpTime = 86400
)

func Run() {
	db := InitRedis()
	repo := repository.NewRepositoryRedis(db)
	service := service2.NewService(repo)
	hndl := handler.NewHandler(service)
	routers := handler.InitRoutes(hndl)

	srv := &http.Server{
		Addr:         serverUrl,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routers, // Pass our instance of gorilla/mux in.
	}
	log.Println("Server started at " + serverUrl)

	go checkUpdate(hndl) // goroutine to refill db one a day

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	rdb.SetNX(context.Background(), "id_counter", 0, 0)
	rdb.SetNX(context.Background(), "timeRefillDB", 0, 0)
	return rdb
}

func checkUpdate(h handler.RefillDB) {
	var bytes []byte
	for {
		expTime, err := h.GetExpTimeDb()
		if err != nil {
			log.Fatal(err)
		}

		if expTime >= dbExpTime {
			bytes, err = h.GetAPIData()
			if err != nil {
				log.Fatal(err)
			}
			err := refillDB(h, bytes)
			if err != nil {
				log.Fatal(err)
			}
			h.FreshExpTimeDb()
			log.Println("REFILLING DATABASE")
		}
		h.IncrExpTimeDb()
		time.Sleep(time.Second * 1)

	}
}

func refillDB(s handler.RefillDB, bytes []byte) error {
	var input []models.TaxiData
	err := json_to_struct.Parse(bytes, &input)
	if err != nil {
		log.Fatal(err.Error())
	}
	s.FlushDB()
	err = s.FillDB(input)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
