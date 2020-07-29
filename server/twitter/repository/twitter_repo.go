package repository

import (
	"github.com/ohatakky/dashboard/server/pkg/twitter"
)

type TwitterRepository interface {
	Tweets() ([]twitter.Tweet, error)
}

type twitterRepository struct {
	twitter *twitter.Client
}

func NewTwitterRepository(tc *twitter.Client) TwitterRepository {
	return &twitterRepository{
		twitter: tc,
	}
}

func (repo *twitterRepository) Tweets() ([]twitter.Tweet, error) {
	return repo.twitter.Tweets()
}
