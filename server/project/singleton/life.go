package singleton

import (
	"log"
	"sync"

	"github.com/ohatakky/dashboard/server/pkg/life"
	"github.com/ohatakky/dashboard/server/project/configs"
)

var Records []life.Record

var onceLife sync.Once

func InitLife() {
	client := life.NewClient(configs.E.Life.SpreadsheetID, configs.E.Life.SheetName)
	onceLife.Do(func() {
		records, err := client.Records()
		if err != nil {
			log.Fatal(err)
		}
		Records = records
	})
}
