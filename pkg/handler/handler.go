package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"taxiTestTask/pkg/service"
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
	router.HandleFunc("/", testF).Methods(http.MethodGet)
	router.HandleFunc("/api/get-id/{id:[0-9]+}", h.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/get-gid/{id:[0-9]+}", h.GetByGlobalId).Methods(http.MethodGet)
	router.HandleFunc("/api/del-id/{id:[0-9]+}", h.DeleteID).Methods(http.MethodDelete)
	router.HandleFunc("/api/del-gid/{id:[0-9]+}", h.DeleteGID).Methods(http.MethodDelete)
	return router
}

func testF(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world!"))
	if err != nil {
		log.Fatal(err)
	}
}
