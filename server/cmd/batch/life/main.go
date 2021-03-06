package main

import (
	"fmt"
	"log"

	"github.com/ohatakky/dashboard/server/pkg/life"
	"github.com/ohatakky/dashboard/server/project/configs"
)

func main() {
	configs.InitConfigs()

	{
		client := life.NewClient(configs.E.Life.SpreadsheetID, configs.E.Life.SheetName)
		res, err := client.Records()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
}
