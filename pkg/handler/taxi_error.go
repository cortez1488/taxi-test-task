package handler

import (
	"encoding/json"
	"net/http"
	"taxiTestTask/pkg/service_errors"
)

var (
	noDataJSON       []byte
	unableToSaveJSON []byte
)

func init() {
	noDataJSON, _ = json.Marshal(map[string]string{"error": "no data for id"})
	unableToSaveJSON, _ = json.Marshal(map[string]string{"error": "unable to save"})
}

func (h *taxiHandler) handleError(w *http.ResponseWriter, err error) {
	switch err {
	case service_errors.ErrNoData:
		(*w).Write(noDataJSON)
	case service_errors.ErrUnableToSave:
		(*w).Write(unableToSaveJSON)
	}
}
