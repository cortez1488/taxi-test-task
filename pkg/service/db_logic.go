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

type DBLogicService struct {
	repo repository.DBLogic
}

func NewDBLogicService(DBLogic repository.DBLogic) *DBLogicService {
	return &DBLogicService{repo: DBLogic}
}

func (s *DBLogicService) RefillDB() error {
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
}

func refillDB(s *DBLogicService) error {
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
