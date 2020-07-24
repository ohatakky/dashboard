package life

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ohatakky/dashboard/server/project/configs"
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

type Weather string

const (
	Sunny  = Weather("晴れ")
	Cloudy = Weather("曇り")
	Rainy  = Weather("雨")
)

const (
	dateLayout = "2006/01/02"
	timeLayout = "15:04:05"
)

type Record struct {
	Date        TimeNull      // 日付
	Condition   int           // 調子[10]
	Rising      time.Time     // 起床
	LightOff    time.Time     // 消灯
	Sleep       time.Duration // 睡眠
	Nap         time.Duration // 仮眠
	Meals       int           // 食事の回数
	Fullness    int           // 満腹感[10]
	Motion      time.Duration // 運動
	Hunting     time.Duration // 狩り
	Devotion    time.Duration // 精進
	Hobby       time.Duration // 趣味
	Meaningless time.Duration // 無意味な時間
	Weather     Weather       // 天気
	S           int           // S
	Vitamin     bool          // ビタミン剤
}

func (c *Client) downloads() ([]byte, error) {
	uFmt := "https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s"
	u := fmt.Sprintf(uFmt, configs.E.Life.SpreadsheetID, configs.E.Life.SheetName)
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

// todo
type TimeNull struct {
	IsNull bool
	time.Time
}

type DurationNull struct {
	IsNull bool
	time.Duration
}

func (c *Client) Records() ([]Record, error) {
	b, err := c.downloads()
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(strings.NewReader(string(b)))
	_, err = r.Read()
	if err != nil {
		return nil, err
	}
	res := make([]Record, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// mappings to struct
		tmp := Record{}
		for i, v := range record {
			var null bool
			if v == "" {
				null = true
			}
			switch i {
			case 0:
				if null {
					tmp.Date.IsNull = true
				}
				tmp.Date.Time, err = time.Parse(dateLayout, v)
				if err != nil {
					return nil, err
				}
			case 1:
				tmp.Condition, err = strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
			case 2:
				tmp.Rising, err = time.Parse(timeLayout, v)
				if err != nil {
					return nil, err
				}
			case 3:
				tmp.LightOff, err = time.Parse(timeLayout, v)
				if err != nil {
					return nil, err
				}
			case 4:
				tmp.Sleep, err = time.ParseDuration(v)
				if err != nil {
					return nil, err
				}
			}
		}
		res = append(res, tmp)
	}
	return res, nil
}
