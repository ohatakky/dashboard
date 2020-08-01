package http

import (
	"encoding/json"
	"net/http"

	"github.com/ohatakky/dashboard/server/atcoder/usecase"
	"github.com/ohatakky/dashboard/server/project/singleton"
)

type HttpAtcoderHandler struct {
	atcoderUC usecase.AtcoderUsecase
}

func NewHttpAtcoderHandler(mux *http.ServeMux, uc usecase.AtcoderUsecase) {
	h := &HttpAtcoderHandler{
		atcoderUC: uc,
	}

	mux.HandleFunc("/atcoder", h.Submissions)
	mux.HandleFunc("/atcoder/latest", h.SubmissionsLatest)
}

func (h *HttpAtcoderHandler) Submissions(w http.ResponseWriter, _ *http.Request) {
	res, err := h.atcoderUC.SubmissionsDaily(singleton.Submissions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (h *HttpAtcoderHandler) SubmissionsLatest(w http.ResponseWriter, _ *http.Request) {
	res, err := h.atcoderUC.Submissions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
