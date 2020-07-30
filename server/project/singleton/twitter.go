package singleton

import (
	"log"
	"sync"

	"github.com/ohatakky/dashboard/server/pkg/twitter"
	"github.com/ohatakky/dashboard/server/project/configs"
)

var Tweets []twitter.Tweet

var once sync.Once

func InitTwitter() {
	client := twitter.NewClient(configs.E.Twitter.SpreadsheetID, configs.E.Twitter.SheetName)
	once.Do(func() {
		tweets, err := client.Tweets()
		if err != nil {
			log.Fatal(err)
		}
		Tweets = tweets
	})
}
