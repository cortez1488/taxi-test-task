package handler

import (
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
	fmt.Println(r.URL.Path)
	w.Write([]byte("get by id"))
}
func (h *taxiHandler) GetByGlobalId(w http.ResponseWriter, r *http.Request) {

}
func (h *taxiHandler) DeleteID(w http.ResponseWriter, r *http.Request) {

}

func (h *taxiHandler) DeleteGID(w http.ResponseWriter, r *http.Request) {

}
