package http

import (
	"encoding/json"
	"net/http"

	"github.com/ohatakky/dashboard/server/note/usecase"
	"github.com/ohatakky/dashboard/server/project/singleton"
)

type HttpNoteHandler struct {
	noteUC usecase.NoteUsecase
}

func NewHttpNoteHandler(mux *http.ServeMux, uc usecase.NoteUsecase) {
	h := &HttpNoteHandler{
		noteUC: uc,
	}

	mux.HandleFunc("/note", h.Posts)
}

func (h *HttpNoteHandler) Posts(w http.ResponseWriter, _ *http.Request) {
	res, err := h.noteUC.PostsDaily(singleton.Posts)
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
