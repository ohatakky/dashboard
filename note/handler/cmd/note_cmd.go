package cmd

import (
	"fmt"
	"log"

	"github.com/ohatakky/dashboard/note/usecase"
)

type CmdNoteHandler struct {
	noteUC usecase.NoteUsecase
}

func NewCmdNoteHandler(uc usecase.NoteUsecase) *CmdNoteHandler {
	return &CmdNoteHandler{
		noteUC: uc,
	}
}

func (h *CmdNoteHandler) GetPosts() {
	res, err := h.noteUC.GetPosts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
