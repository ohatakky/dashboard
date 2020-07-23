package main

import (
	"github.com/ohatakky/dashboard/atcoder/handler/cmd"
	"github.com/ohatakky/dashboard/atcoder/repository"
	"github.com/ohatakky/dashboard/atcoder/usecase"
	"github.com/ohatakky/dashboard/pkg/atcoder"
	"github.com/ohatakky/dashboard/project/configs"
)

func main() {
	// todo: 共通処理に切り出す
	configs.InitConfigs()

	{
		client := atcoder.NewClient(configs.E.Atcoder.User)
		repo := repository.NewAtcoderRepository(client)
		uc := usecase.NewAtcoderUsecase(repo)
		cmdHandler := cmd.NewCmdAtcoderHandler(uc)
		cmdHandler.Submissions()
	}
}
