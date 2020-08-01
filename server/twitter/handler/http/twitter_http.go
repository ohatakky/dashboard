package http

import (
	"encoding/json"
	"net/http"

	"github.com/ohatakky/dashboard/server/project/singleton"
	"github.com/ohatakky/dashboard/server/twitter/usecase"
)

type HttpTwitterHandler struct {
	twitterUC usecase.TwitterUsecase
}

func NewHttpTwitterHandler(mux *http.ServeMux, uc usecase.TwitterUsecase) {
	h := &HttpTwitterHandler{
		twitterUC: uc,
	}

	mux.HandleFunc("/twitter", h.Tweets)
	mux.HandleFunc("/twitter/latest", h.TweetsLatest)
}

func (h *HttpTwitterHandler) Tweets(w http.ResponseWriter, _ *http.Request) {
	b, err := json.Marshal(singleton.Tweets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (h *HttpTwitterHandler) TweetsLatest(w http.ResponseWriter, _ *http.Request) {
	res, err := h.twitterUC.Tweets()
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
