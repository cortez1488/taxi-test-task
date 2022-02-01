package db_update

import (
	"log"
	"taxiTestTask/internal/json_to_struct"
	reqToAPI "taxiTestTask/internal/reqToAPI/JSON"
	"taxiTestTask/models"
	"taxiTestTask/pkg/repository"
)

func RefillDB(repo *repository.Repository) error {
	json, err := reqToAPI.GetJSONFromAPIRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	var input []models.TaxiData
	err = json_to_struct.Parse(json, &input)
	if err != nil {
		log.Fatal(err.Error())
	}
	repo.FlushDB()
	err = repo.FillDB(&input)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
