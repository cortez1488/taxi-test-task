package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"taxiTestTask/pkg/service"
)

type taxiHandler struct {
	service *service.Service
}

func newTaxiHandler(service *service.Service) *taxiHandler {
	return &taxiHandler{service: service}
}

func (h *taxiHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	data, err := h.service.GetById(id)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}

}
func (h *taxiHandler) GetByGlobalId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gid, err := strconv.ParseInt(vars["gid"], 10, 64)

	data, err := h.service.GetByGlobalId(gid)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
func (h *taxiHandler) DeleteID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	succ, err := h.service.DeleteID(id)
	if err != nil {
		log.Fatal(err)
	}

	if succ == 1 {
		w.WriteHeader(http.StatusNoContent)
	} else if succ == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *taxiHandler) DeleteGID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gid, err := strconv.ParseInt(vars["gid"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	succ, err := h.service.DeleteGID(gid)
	if err != nil {
		log.Fatal(err)
	}
	if succ == 1 {
		w.WriteHeader(http.StatusNoContent)
	} else if succ == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
}
