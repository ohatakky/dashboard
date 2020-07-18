package usecase

import (
	"github.com/ohatakky/dashboard/atcoder"
	"github.com/ohatakky/dashboard/atcoder/repository"
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
