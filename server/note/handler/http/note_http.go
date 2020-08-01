package http

import (
	"encoding/json"
	"net/http"

	"github.com/ohatakky/dashboard/server/project/singleton"
)

type HttpNoteHandler struct{}

func NewHttpNoteHandler(mux *http.ServeMux) {
	h := &HttpNoteHandler{}

	mux.HandleFunc("/note", h.Posts)
}

func (h *HttpNoteHandler) Posts(w http.ResponseWriter, _ *http.Request) {
	b, err := json.Marshal(singleton.Posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
