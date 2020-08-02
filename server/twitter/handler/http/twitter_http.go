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
}

func (h *HttpTwitterHandler) Tweets(w http.ResponseWriter, _ *http.Request) {
	res, err := h.twitterUC.TweetsDaily(singleton.Tweets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
