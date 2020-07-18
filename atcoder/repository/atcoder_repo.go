package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/ohatakky/dashboard/atcoder"
	"github.com/ohatakky/dashboard/project/configs"
)

type AtcoderRepository interface {
	Submissions() ([]atcoder.Submission, error)
}

type atcoderRepository struct {
}

func NewAtcoderRepository() AtcoderRepository {
	return &atcoderRepository{}
}

func (repo *atcoderRepository) Submissions() ([]atcoder.Submission, error) {
	u, err := url.Parse(atcoder.Endpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, atcoder.SubmissionPath)
	params := url.Values{}
	params.Add("user", configs.E.Atcoder.User)
	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := []atcoder.Submission{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
