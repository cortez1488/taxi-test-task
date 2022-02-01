package service

import (
	"log"
	"taxiTestTask/internal/json_to_struct"
	reqToAPI "taxiTestTask/internal/reqToAPI/JSON"
	"taxiTestTask/models"
	"taxiTestTask/pkg/repository"
)

type taxiService struct {
	repo repository.TaxiParking
}

func newTaxiService(repo *repository.TaxiParking) *taxiService {
	return &taxiService{repo: *repo}
}

func (s *taxiService) GetById(id int) (*models.TaxiData, error) {
	return s.repo.GetById(id)
}

func (s *taxiService) GetByGlobalId(id int64) (*models.TaxiData, error) {
	return s.repo.GetByGlobalId(id)
}

func (s *taxiService) DeleteID(id int) (int64, error) {
	return s.repo.DeleteID(id)
}

func (s *taxiService) DeleteGID(id int64) (int64, error) {
	return s.repo.DeleteGID(id)
}

func (s *taxiService) RefillDB(*[]models.TaxiData) error {
	json, err := reqToAPI.GetJSONFromAPIRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	var input []models.TaxiData
	err = json_to_struct.Parse(json, &input)
	if err != nil {
		log.Fatal(err.Error())
	}
	s.repo.FlushDB()
	err = s.repo.FillDB(&input)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
