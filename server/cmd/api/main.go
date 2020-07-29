package main

import (
	"log"
	"net/http"

	acHandler "github.com/ohatakky/dashboard/server/atcoder/handler/http"
	acRepo "github.com/ohatakky/dashboard/server/atcoder/repository"
	acUsecase "github.com/ohatakky/dashboard/server/atcoder/usecase"

	twHandler "github.com/ohatakky/dashboard/server/twitter/handler/http"
	twRepo "github.com/ohatakky/dashboard/server/twitter/repository"
	twUsecase "github.com/ohatakky/dashboard/server/twitter/usecase"

	"github.com/ohatakky/dashboard/server/pkg/atcoder"
	"github.com/ohatakky/dashboard/server/pkg/twitter"
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
	{
		client := twitter.NewClient(configs.E.Twitter.SpreadsheetID, configs.E.Twitter.SheetName)
		repo := twRepo.NewTwitterRepository(client)
		uc := twUsecase.NewTwitterUsecase(repo)
		twHandler.NewHttpTwitterHandler(mux, uc)
	}

	log.Fatal(http.ListenAndServe(":8080", mux))
}
