package cmd

import (
	"fmt"
	"log"

	"github.com/ohatakky/dashboard/server/atcoder/usecase"
)

type CmdAtcoderHandler struct {
	atcoderUC usecase.AtcoderUsecase
}

func NewCmdAtcoderHandler(uc usecase.AtcoderUsecase) *CmdAtcoderHandler {
	return &CmdAtcoderHandler{
		atcoderUC: uc,
	}
}

func (h *CmdAtcoderHandler) Submissions() {
	res, err := h.atcoderUC.Submissions()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
