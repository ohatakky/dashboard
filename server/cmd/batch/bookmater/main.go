package main

import (
	"fmt"
	"log"

	"github.com/ohatakky/dashboard/server/pkg/bookmater"
	"github.com/ohatakky/dashboard/server/project/configs"
)

func main() {
	// todo: 共通処理に切り出す
	configs.InitConfigs()

	{
		client := bookmater.NewClient(configs.E.Bookmater.User)
		res, err := client.GetReviews()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
}
