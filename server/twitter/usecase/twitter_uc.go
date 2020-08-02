package usecase

import (
	"sort"
	"time"

	"github.com/ohatakky/dashboard/server/pkg/twitter"
)

type TwitterUsecase interface {
	TweetsDaily([]twitter.Tweet) ([]TweetsDailyResponse, error)
}

type twitterUsecase struct {
}

func NewTwitterUsecase() TwitterUsecase {
	return &twitterUsecase{}
}

func (uc *twitterUsecase) TweetsDaily(tweets []twitter.Tweet) ([]TweetsDailyResponse, error) {
	layout := "2006-01-02"
	aggregates := make(map[string]int)
	from := tweets[0].CreatedAt
	to := time.Now()

	for d := from; !d.After(to); d = d.AddDate(0, 0, 1) {
		dateFmt := d.Format(layout)
		aggregates[dateFmt] = 0
	}

	for _, s := range tweets {
		dateFmt := s.CreatedAt.Format(layout)
		aggregates[dateFmt]++
	}

	res := []TweetsDailyResponse{}
	for daily, count := range aggregates {
		res = append(res, TweetsDailyResponse{
			Date:  daily,
			Count: count,
		})
	}
	sort.Slice(res, func(i, j int) bool { return res[i].Date < res[j].Date })

	return res, nil
}
