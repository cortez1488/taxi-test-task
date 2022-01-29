package main

import (
	"fmt"
	"log"
	"taxiTestTask/internal/json_to_struct"
	reqToAPI "taxiTestTask/internal/reqToAPI/JSON"
	"taxiTestTask/models"
)

func main() {

	json, err := reqToAPI.GetJSONFromAPIRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	var input []models.TaxiRawData
	err = json_to_struct.ParseBodyJson(json, &input)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(input)
}
