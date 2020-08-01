package singleton

import (
	"log"
	"sync"

	"github.com/ohatakky/dashboard/server/pkg/atcoder"
	"github.com/ohatakky/dashboard/server/project/configs"
)

var Submissions []atcoder.Submission

var onceAtcoder sync.Once

func InitAtcoder() {
	client := atcoder.NewClient(configs.E.Atcoder.User)
	onceAtcoder.Do(func() {
		submissions, err := client.Submissions()
		if err != nil {
			log.Fatal(err)
		}
		Submissions = submissions
	})
}
