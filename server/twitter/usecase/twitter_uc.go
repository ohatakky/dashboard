package usecase

import (
	"github.com/ohatakky/dashboard/server/pkg/twitter"
	"github.com/ohatakky/dashboard/server/twitter/repository"
)

type TwitterUsecase interface {
	Tweets() ([]twitter.Tweet, error)
}

type twitterUsecase struct {
	twitterRepo repository.TwitterRepository
}

func NewTwitterUsecase(tr repository.TwitterRepository) TwitterUsecase {
	return &twitterUsecase{
		twitterRepo: tr,
	}
}

func (uc *twitterUsecase) Tweets() ([]twitter.Tweet, error) {
	return uc.twitterRepo.Tweets()
}
