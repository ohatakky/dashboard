package usecase

import (
	"github.com/ohatakky/dashboard/server/atcoder/repository"
	"github.com/ohatakky/dashboard/server/pkg/atcoder"
)

type AtcoderUsecase interface {
	Submissions() ([]atcoder.Submission, error)
}

type atcoderUsecase struct {
	atcoderRepo repository.AtcoderRepository
}

func NewAtcoderUsecase(ar repository.AtcoderRepository) AtcoderUsecase {
	return &atcoderUsecase{
		atcoderRepo: ar,
	}
}

func (uc *atcoderUsecase) Submissions() ([]atcoder.Submission, error) {
	return uc.atcoderRepo.Submissions()
}
