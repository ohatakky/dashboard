package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"

	acHandler "github.com/ohatakky/dashboard/server/atcoder/handler/http"
	acRepo "github.com/ohatakky/dashboard/server/atcoder/repository"
	acUsecase "github.com/ohatakky/dashboard/server/atcoder/usecase"
	"github.com/ohatakky/dashboard/server/pkg/atcoder"
	"github.com/ohatakky/dashboard/server/project/configs"
	"github.com/ohatakky/dashboard/server/project/singleton"
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
	// {
	// 	client := twitter.NewClient(configs.E.Twitter.SpreadsheetID, configs.E.Twitter.SheetName)
	// 	repo := twRepo.NewTwitterRepository(client)
	// 	uc := twUsecase.NewTwitterUsecase(repo)
	// 	twHandler.NewHttpTwitterHandler(mux, uc)
	// }
	{
		singleton.InitTwitter()

		mux.HandleFunc("/twitter", func(w http.ResponseWriter, _ *http.Request) {
			b, err := json.Marshal(singleton.Tweets)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		})
	}

	c := cors.New(cors.Options{
		// AllowedOrigins: []string{"http://localhost:3001"},
		AllowedMethods: []string{"GET"},
	})
	handler := c.Handler(mux)

	log.Println("running...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
