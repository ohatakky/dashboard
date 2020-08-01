package http

import (
	"encoding/json"
	"net/http"

	"github.com/ohatakky/dashboard/server/project/singleton"
)

type HttpBookmaterHandler struct{}

func NewHttpBookmaterHandler(mux *http.ServeMux) {
	h := &HttpBookmaterHandler{}

	mux.HandleFunc("/bookmater", h.Reviews)
}

func (h *HttpBookmaterHandler) Reviews(w http.ResponseWriter, _ *http.Request) {
	b, err := json.Marshal(singleton.Reviews)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
