package app

import (
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"taxiTestTask/pkg/handler"
	"taxiTestTask/pkg/repository"
	service2 "taxiTestTask/pkg/service"
	"time"
)

const (
	serverUrl = "localhost:8000"
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

	go refillDb(service) // goroutine to refill db one a day

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func refillDb(service *service2.Service) {
	err := service.RefillDB()
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
	return rdb
}
