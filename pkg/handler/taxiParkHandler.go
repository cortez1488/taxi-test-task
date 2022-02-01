package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"taxiTestTask/pkg/service"
)

type taxiHandler struct {
	service *service.Service
}

func newTaxiHandler(service *service.Service) *taxiHandler {
	return &taxiHandler{service: service}
}

func (h *taxiHandler) GetById(w http.ResponseWriter, r *http.Request) {
	path := ([]byte(r.URL.Path))
	id := bytes.TrimPrefix(path, []byte(routeGetId))
	w.Write([]byte(fmt.Sprintf("Your id is: %s", id)))
}
func (h *taxiHandler) GetByGlobalId(w http.ResponseWriter, r *http.Request) {

}
func (h *taxiHandler) DeleteID(w http.ResponseWriter, r *http.Request) {

}

func (h *taxiHandler) DeleteGID(w http.ResponseWriter, r *http.Request) {

}
