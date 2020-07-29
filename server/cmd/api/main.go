package main

import (
	"log"
	"net/http"

	acHandler "github.com/ohatakky/dashboard/server/atcoder/handler/http"
	acRepo "github.com/ohatakky/dashboard/server/atcoder/repository"
	acUsecase "github.com/ohatakky/dashboard/server/atcoder/usecase"
	"github.com/ohatakky/dashboard/server/pkg/atcoder"
	"github.com/ohatakky/dashboard/server/project/configs"
)

func main() {
	configs.InitConfigs()
	mux := http.NewServeMux()
	{
		client := atcoder.NewClient(configs.E.Atcoder.User)
		repo := acRepo.NewAtcoderRepository(client)
		uc := acUsecase.NewAtcoderUsecase(repo)
		acHandler.NewHttpAtcoderHandler(mux, uc)
	}
	log.Fatal(http.ListenAndServe(":8080", mux))
}
