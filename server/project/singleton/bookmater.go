package singleton

import (
	"log"
	"sync"

	"github.com/ohatakky/dashboard/server/pkg/bookmater"
	"github.com/ohatakky/dashboard/server/project/configs"
)

var Reviews *bookmater.Reviews

var onceBookmater sync.Once

func InitBookmater() {
	client := bookmater.NewClient(configs.E.Bookmater.User)
	onceBookmater.Do(func() {
		reviews, err := client.GetReviews()
		if err != nil {
			log.Fatal(err)
		}
		Reviews = reviews
	})
}
