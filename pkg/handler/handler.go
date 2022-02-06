package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"taxiTestTask/models"
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
	RefillDB
}

type TaxiParking interface {
	GetById(http.ResponseWriter, *http.Request)
	GetByGlobalId(http.ResponseWriter, *http.Request)
	DeleteID(http.ResponseWriter, *http.Request)
	DeleteGID(http.ResponseWriter, *http.Request)
}

type RefillDB interface {
	GetAPIData() ([]byte, error)

	FillDB([]models.TaxiData) error
	FlushDB()
	GetExpTimeDb() (int, error)
	FreshExpTimeDb()
	IncrExpTimeDb()
}

func NewHandler(service *service.Service) Handler {
	return Handler{TaxiParking: newTaxiHandler(service),
		RefillDB: newRefillDBHandler(service.DBLogic)}
}

func InitRoutes(h Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(routeGetId+"{id:[0-9]+}", h.GetById).Methods(http.MethodGet)
	router.HandleFunc(routeGetGid+"{gid:[0-9]+}", h.GetByGlobalId).Methods(http.MethodGet)
	router.HandleFunc(routeDelId+"{id:[0-9]+}", h.DeleteID).Methods(http.MethodDelete)
	router.HandleFunc(routeDelGid+"{gid:[0-9]+}", h.DeleteGID).Methods(http.MethodDelete)
	return router
}
