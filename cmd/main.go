package main

import (
	"fmt"
	"log"
	reqToAPI "taxiTestTask/internal/reqToAPI/JSON"
	"taxiTestTask/models"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var input []models.TaxiRawData
	err := reqToAPI.RequestJSON(&input)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(input)
}
