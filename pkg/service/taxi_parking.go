package service

import (
	"errors"
	"log"
	"taxiTestTask/internal/json_to_struct"
	reqToAPI "taxiTestTask/internal/reqToAPI/JSON"
	"taxiTestTask/models"
	"taxiTestTask/pkg/repository"
	time2 "time"
)

const (
	dbExpTime = 86400
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

func (s *taxiService) RefillDB() error {
	for {
		time, err := s.repo.GetExpTimeDb()
		if err != nil {
			return errors.New("s.repo.GetExpTimeDb() :" + err.Error())
		}

		if time >= dbExpTime {
			err := refillDB(s)
			if err != nil {
				return errors.New("func refillDB(s) :" + err.Error())
			}
			s.repo.FreshExpTimeDb()
			log.Println("REFILLING DATABASE")
		}
		s.repo.IncrExpTimeDb()
		time2.Sleep(time2.Second * 1)

	}
	return nil
}

func refillDB(s *taxiService) error {
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
