package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	acHttp "github.com/ohatakky/dashboard/server/atcoder/handler/http"
	acRepo "github.com/ohatakky/dashboard/server/atcoder/repository"
	acUsecase "github.com/ohatakky/dashboard/server/atcoder/usecase"
	boHttp "github.com/ohatakky/dashboard/server/bookmater/handler/http"
	liHttp "github.com/ohatakky/dashboard/server/life/handler/http"
	noHttp "github.com/ohatakky/dashboard/server/note/handler/http"
	twHttp "github.com/ohatakky/dashboard/server/twitter/handler/http"
	twRepo "github.com/ohatakky/dashboard/server/twitter/repository"
	twUsecase "github.com/ohatakky/dashboard/server/twitter/usecase"

	"github.com/ohatakky/dashboard/server/pkg/atcoder"
	"github.com/ohatakky/dashboard/server/pkg/twitter"
	"github.com/ohatakky/dashboard/server/project/configs"
	"github.com/ohatakky/dashboard/server/project/singleton"
)

func main() {
	configs.InitConfigs()

	mux := http.NewServeMux()
	{
		singleton.InitAtcoder()
		client := atcoder.NewClient(configs.E.Atcoder.User)
		repo := acRepo.NewAtcoderRepository(client)
		uc := acUsecase.NewAtcoderUsecase(repo)
		acHttp.NewHttpAtcoderHandler(mux, uc)
	}
	{
		singleton.InitBookmater()
		boHttp.NewHttpBookmaterHandler(mux)
	}
	{
		singleton.InitLife()
		liHttp.NewHttpLifeHandler(mux)
	}
	{
		singleton.InitNote()
		noHttp.NewHttpNoteHandler(mux)
	}
	{
		singleton.InitTwitter()
		client := twitter.NewClient(configs.E.Twitter.SpreadsheetID, configs.E.Twitter.SheetName)
		repo := twRepo.NewTwitterRepository(client)
		uc := twUsecase.NewTwitterUsecase(repo)
		twHttp.NewHttpTwitterHandler(mux, uc)
	}

	c := cors.New(cors.Options{
		// AllowedOrigins: []string{"http://localhost:3001"},
		AllowedMethods: []string{"GET"},
	})
	handler := c.Handler(mux)

	log.Println("running...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
