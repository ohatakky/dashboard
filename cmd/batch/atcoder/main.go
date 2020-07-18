package main

import (
	"github.com/ohatakky/dashboard/atcoder/handler/cmd"
	"github.com/ohatakky/dashboard/atcoder/repository"
	"github.com/ohatakky/dashboard/atcoder/usecase"
	"github.com/ohatakky/dashboard/project/configs"
)

func main() {
	// todo: 共通処理に切り出す
	configs.InitConfigs()

	{
		atcoderRepo := repository.NewAtcoderRepository()
		atcoderUC := usecase.NewAtcoderUsecase(atcoderRepo)
		atcoderCmdHandler := cmd.NewCmdAtcoderHandler(atcoderUC)
		atcoderCmdHandler.Submissions()
	}
}
