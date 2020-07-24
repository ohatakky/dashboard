package main

import (
	"github.com/ohatakky/dashboard/server/atcoder/handler/cmd"
	"github.com/ohatakky/dashboard/server/atcoder/repository"
	"github.com/ohatakky/dashboard/server/atcoder/usecase"
	"github.com/ohatakky/dashboard/server/pkg/atcoder"
	"github.com/ohatakky/dashboard/server/project/configs"
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
