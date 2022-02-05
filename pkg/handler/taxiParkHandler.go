package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := h.service.GetById(id)
	if err != nil {
		h.handleError(&w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *taxiHandler) GetByGlobalId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gid, err := strconv.ParseInt(vars["gid"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := h.service.GetByGlobalId(gid)
	if err != nil {
		h.handleError(&w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (h *taxiHandler) DeleteID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	succ, err := h.service.DeleteID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	succ, err := h.service.DeleteGID(gid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if succ == 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	} else if succ == 0 && err == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
