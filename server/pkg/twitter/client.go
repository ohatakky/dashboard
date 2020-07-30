package twitter

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	spreadsheetID string
	sheetName     string
}

func NewClient(spreadsheetID, sheetName string) *Client {
	return &Client{
		spreadsheetID: spreadsheetID,
		sheetName:     sheetName,
	}
}

type Tweet struct {
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
}

func (c *Client) downloads() ([]byte, error) {
	uFmt := "https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s"
	u := fmt.Sprintf(uFmt, c.spreadsheetID, c.sheetName)
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) Tweets() ([]Tweet, error) {
	b, err := c.downloads()
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(strings.NewReader(string(b)))
	_, err = r.Read()
	if err != nil {
		return nil, err
	}
	res := make([]Tweet, 0)
	for {
		tweet, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		tmp := Tweet{}
		for i, t := range tweet {
			if t == "" {
				continue
			}
			switch i {
			case 0:
				// tmp.CreatedAt = t
			case 1:
				tmp.Text = t
			}
		}
		res = append(res, tmp)
	}
	return res, nil
}
