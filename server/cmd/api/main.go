package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/ohatakky/dashboard/server/project/configs"
	"github.com/ohatakky/dashboard/server/project/singleton"
)

func main() {
	configs.InitConfigs()

	mux := http.NewServeMux()

	{
		singleton.InitAtcoder()
		mux.HandleFunc("/atcoder", func(w http.ResponseWriter, _ *http.Request) {
			b, err := json.Marshal(singleton.Submissions)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		})
	}
	{
		singleton.InitBookmater()
		mux.HandleFunc("/bookmater", func(w http.ResponseWriter, _ *http.Request) {
			b, err := json.Marshal(singleton.Reviews)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		})
	}
	{
		singleton.InitNote()
		mux.HandleFunc("/note", func(w http.ResponseWriter, _ *http.Request) {
			b, err := json.Marshal(singleton.Posts)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		})
	}
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
