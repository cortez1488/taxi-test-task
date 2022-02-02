package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"taxiTestTask/pkg/service"
)

const (
	routeGetId  = "/api/get-id/"
	routeGetGid = "/api/get-gid/"
	routeDelId  = "/api/del-id/"
	routeDelGid = "/api/del-gid/"
)

type Handler struct {
	TaxiParking
}

type TaxiParking interface {
	GetById(http.ResponseWriter, *http.Request)
	GetByGlobalId(http.ResponseWriter, *http.Request)
	DeleteID(http.ResponseWriter, *http.Request)
	DeleteGID(http.ResponseWriter, *http.Request)
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{TaxiParking: newTaxiHandler(service)}
}

func InitRoutes(h *Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(routeGetId+"{id:[0-9]+}", h.GetById).Methods(http.MethodGet)
	router.HandleFunc(routeGetGid+"{gid:[0-9]+}", h.GetByGlobalId).Methods(http.MethodGet)
	router.HandleFunc(routeDelId+"{id:[0-9]+}", h.DeleteID).Methods(http.MethodDelete)
	router.HandleFunc(routeDelGid+"{gid:[0-9]+}", h.DeleteGID).Methods(http.MethodDelete)
	return router
}
