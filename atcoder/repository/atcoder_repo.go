package repository

import (
	"github.com/ohatakky/dashboard/pkg/atcoder"
)

type AtcoderRepository interface {
	Submissions() ([]atcoder.Submission, error)
}

type atcoderRepository struct {
	atcoder *atcoder.Client
}

func NewAtcoderRepository(ac *atcoder.Client) AtcoderRepository {
	return &atcoderRepository{
		atcoder: ac,
	}
}

func (repo *atcoderRepository) Submissions() ([]atcoder.Submission, error) {
	return repo.atcoder.Submissions()
}
