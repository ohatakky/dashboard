package http

import (
	"encoding/json"
	"net/http"

	"github.com/ohatakky/dashboard/server/project/singleton"
)

type HttpLifeHandler struct{}

func NewHttpLifeHandler(mux *http.ServeMux) {
	h := &HttpLifeHandler{}

	mux.HandleFunc("/life", h.Records)
}

func (h *HttpLifeHandler) Records(w http.ResponseWriter, _ *http.Request) {
	b, err := json.Marshal(singleton.Records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
