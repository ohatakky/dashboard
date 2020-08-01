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

const (
	dateLayout     = "2006/01/02"
	timeLayout     = "15:04:05"
	durationLayout = "1h00m"
)

func dateParser() {

}

func timeParser() {

}

func durationParser() {

}

type Time struct {
	time.Time
	Valid bool
}

type Duration struct {
	time.Duration
	Valid bool
}

type String struct {
	string
	Valid bool
}

type Int struct {
	int
	Valid bool
}

type Float struct {
	float64
	Valid bool
}

type Bool struct {
	bool
	Valid bool
}

type Record struct {
	Date        Time     // 日付
	Condition   Int      // 調子
	Rising      Time     // 起床
	Sleep       Duration // 睡眠
	LightOff    Bool     // 消灯
	Bath        Time     // 風呂
	Fullness    Int      // 満腹感
	Vitamin     Bool     // ビタミン剤
	Weather     String   // 天気
	Hunting     Duration // 狩
	Devotion    Duration // 精進
	Hobby       Duration // 趣味
	Meaningless Duration // 無の時間
	WorkoutW    Duration // ワークアウトW
	WorkoutR    Float    // ワークアウトR
	WorkoutB    Int      // ワークアウトB
	WorkoutE    Int      // ワークアウトE
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

		tmp := Record{}
		for i, v := range record {
			var null bool
			if v == "" {
				null = true
			}
			switch i {
			case 0:
				tmp.Date.Valid = null
				if null {
					continue
				}
				tmp.Date.Time, err = time.Parse(dateLayout, v)
				if err != nil {
					return nil, err
				}
			case 1:
				tmp.Condition.Valid = null
				if null {
					continue
				}
				tmp.Condition.int, err = strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
			case 2:
				tmp.Rising.Valid = null
				if null {
					continue
				}
				tmp.Rising.Time, err = time.Parse(timeLayout, v)
				if err != nil {
					return nil, err
				}
			case 3:
				tmp.Sleep.Valid = null
				if null {
					continue
				}
				tmp.Sleep.Duration, err = time.ParseDuration(v)
				if err != nil {
					return nil, err
				}
			case 4:
				tmp.LightOff.Valid = null
				if null {
					continue
				}
				if v == "Yes" {
					tmp.LightOff.bool = true
				} else if v == "No" {
					tmp.LightOff.bool = false
				}
			}
		}
		res = append(res, tmp)
	}
	return res, nil
}
