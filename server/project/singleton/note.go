package singleton

import (
	"log"
	"sync"

	"github.com/ohatakky/dashboard/server/pkg/note"
	"github.com/ohatakky/dashboard/server/project/configs"
)

var Posts *note.Posts

var onceNote sync.Once

func InitNote() {
	client := note.NewClient(configs.E.Note.User)
	onceNote.Do(func() {
		posts, err := client.GetPosts()
		if err != nil {
			log.Fatal(err)
		}
		Posts = posts
	})
}
