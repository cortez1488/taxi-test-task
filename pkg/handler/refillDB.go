package handler

import (
	"fmt"
	"io"
	"net/http"
	"taxiTestTask/models"
	"taxiTestTask/pkg/service"
)

const (
	apiKey = "85a31f5108e65a7e9bbd6c0ade6ae33b"
	uri    = "https://apidata.mos.ru/v1/datasets/621/features?api_key=%s"
)

type RefillDBHandler struct {
	service service.DBLogic
}

func newRefillDBHandler(service service.DBLogic) *RefillDBHandler {
	return &RefillDBHandler{service: service}
}

func (s *RefillDBHandler) FillDB(data []models.TaxiData) error {
	return s.service.FillDB(data)
}

func (s *RefillDBHandler) FlushDB() {
	s.service.FlushDB()
}

func (s *RefillDBHandler) GetExpTimeDb() (int, error) {
	return s.service.GetExpTimeDb()
}

func (s *RefillDBHandler) IncrExpTimeDb() {
	s.service.IncrExpTimeDb()
}

func (s *RefillDBHandler) FreshExpTimeDb() {
	s.service.FreshExpTimeDb()
}

func (s *RefillDBHandler) GetAPIData() ([]byte, error) {
	url := fmt.Sprintf(uri, apiKey)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}
