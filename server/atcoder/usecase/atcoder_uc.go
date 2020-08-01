package usecase

import (
	"sort"
	"time"

	"github.com/ohatakky/dashboard/server/atcoder/repository"
	"github.com/ohatakky/dashboard/server/pkg/atcoder"
)

type AtcoderUsecase interface {
	Submissions() ([]atcoder.Submission, error)
	SubmissionsDaily([]atcoder.Submission) ([]SubmissionsDailyResponse, error)
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

func (uc *atcoderUsecase) SubmissionsDaily(submissions []atcoder.Submission) ([]SubmissionsDailyResponse, error) {
	layout := "2006-01-02"
	aggregates := make(map[string]int)
	from := time.Unix(int64(submissions[0].EpochSecond), 0)
	to := time.Unix(int64(submissions[len(submissions)-1].EpochSecond), 0)

	for d := from; d.Year() <= to.Year() && d.Month() <= to.Month() && d.Day() <= to.Day(); d = d.AddDate(0, 0, 1) {
		dateFmt := d.Format(layout)
		aggregates[dateFmt] = 0
	}

	for _, s := range submissions {
		dateFmt := time.Unix(int64(s.EpochSecond), 0).Format(layout)
		aggregates[dateFmt]++
	}

	res := []SubmissionsDailyResponse{}
	for daily, count := range aggregates {
		res = append(res, SubmissionsDailyResponse{
			Date:  daily,
			Count: count,
		})
	}
	sort.Slice(res, func(i, j int) bool { return res[i].Date < res[j].Date })

	return res, nil
}
