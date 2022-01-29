package app

import (
	"fmt"
	"log"
	"taxiTestTask/internal/json_to_struct"
	reqToAPI "taxiTestTask/internal/reqToAPI/JSON"
	"taxiTestTask/models"
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
	fmt.Println(input)
}
