package atcoder

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const (
	endpoint       = "https://kenkoooo.com/atcoder/atcoder-api"
	submissionPath = "results"
)

type Submission struct {
	ID            int     `json:"id"`
	EpochSecond   int     `json:"epoch_second"`
	ProblemID     string  `json:"problem_id"`
	ContestID     string  `json:"contest_id"`
	UserID        string  `json:"user_id"`
	Language      string  `json:"language"`
	Point         float64 `json:"point"`
	Length        int     `json:"length"`
	Result        string  `json:"result"`
	ExecutionTime int     `json:"execution_time"`
}

type Client struct {
	user string
}

func NewClient(user string) *Client {
	return &Client{
		user: user,
	}
}

func (c *Client) Submissions() ([]Submission, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, submissionPath)
	params := url.Values{}
	params.Add("user", c.user)
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

	res := []Submission{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
