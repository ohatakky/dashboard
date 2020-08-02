package note

import "time"

type Post struct {
	Name      string    `json:"name"`
	Body      string    `json:"body"`
	PublishAt time.Time `json:"publishAt"`
}
