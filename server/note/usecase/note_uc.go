package usecase

import (
	"sort"
	"time"

	_note "github.com/ohatakky/dashboard/server/note"
	"github.com/ohatakky/dashboard/server/pkg/note"
)

type NoteUsecase interface {
	PostsDaily([]note.Posts) ([]PostsDailyResponse, error)
}

type noteUsecase struct{}

func NewNoteUsecase() NoteUsecase {
	return &noteUsecase{}
}

func (uc *noteUsecase) PostsDaily(postsPager []note.Posts) ([]PostsDailyResponse, error) {
	noteLayout := "2006/01/02 15:04"
	tmp := make([]_note.Post, 0)
	for _, posts := range postsPager {
		for _, post := range posts.Data.Contents {
			pub, err := time.Parse(noteLayout, post.PublishAt)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, _note.Post{
				Name:      post.Name,
				Body:      post.Body,
				PublishAt: pub,
			})
		}
	}

	layout := "2006-01-02"
	aggregates := make(map[string]int)
	from := tmp[len(tmp)-1].PublishAt
	to := tmp[0].PublishAt
	for d := from; d.Year() <= to.Year() && d.Month() <= to.Month() && d.Day() <= to.Day(); d = d.AddDate(0, 0, 1) {
		dateFmt := d.Format(layout)
		aggregates[dateFmt] = 0
	}

	for _, s := range tmp {
		dateFmt := s.PublishAt.Format(layout)
		aggregates[dateFmt]++
	}

	res := []PostsDailyResponse{}
	for daily, count := range aggregates {
		res = append(res, PostsDailyResponse{
			Date:  daily,
			Count: count,
		})
	}
	sort.Slice(res, func(i, j int) bool { return res[i].Date < res[j].Date })

	return res, nil
}
