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

type Time struct {
	Time  time.Time `json:"time"`
	Valid bool      `json:"valid"`
}

type Duration struct {
	Duration time.Duration `json:"duration"`
	Valid    bool          `json:"valid"`
}

type String struct {
	String string `json:"string"`
	Valid  bool   `json:"valid"`
}

type Int struct {
	Int   int  `json:"int"`
	Valid bool `json:"valid"`
}

type Float struct {
	Float float64 `json:"float"`
	Valid bool    `json:"valid"`
}

type Bool struct {
	Bool  bool `json:"bool"`
	Valid bool `json:"valid"`
}

type Record struct {
	Date      Time     `json:"date"`      // 日付
	Condition Int      `json:"condition"` // 調子
	Rising    Time     `json:"rising"`    // 起床
	Sleep     Duration `json:"sleep"`     // 睡眠
	LightOff  Bool     `json:"light_off"` // 消灯
	Bath      Time     `json:"bath"`      // 風呂
	Fullness  Int      `json:"fullness"`  // 満腹感
	Vitamin   Bool     `json:"vitamin"`   // ビタミン剤
	Weather   String   `json:"weather"`   // 天気
	Hunting   Duration `json:"hunting"`   // 狩
	Devotion  Duration `json:"devotion"`  // 精進
	Hobby     Duration `json:"hobby"`     // 趣味
	WorkoutW  Duration `json:"workout_w"` // ワークアウトW
	WorkoutR  Float    `json:"workout_r"` // ワークアウトR
	WorkoutB  Int      `json:"workout_b"` // ワークアウトB
	// WorkoutE  Int      `json:"workout_e"` // ワークアウトE
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
				tmp.Condition.Int, err = strconv.Atoi(v)
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
				tmp.LightOff.Bool = v == "Yes"
			case 5:
				tmp.Bath.Valid = null
				if null {
					continue
				}
				tmp.Bath.Time, err = time.Parse(timeLayout, v)
				if err != nil {
					return nil, err
				}
			case 6:
				tmp.Fullness.Valid = null
				if null {
					continue
				}
				tmp.Fullness.Int, err = strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
			case 7:
				tmp.Vitamin.Valid = null
				if null {
					continue
				}
				tmp.Vitamin.Bool = v == "Yes"
			case 8:
				tmp.Weather.Valid = null
				if null {
					continue
				}
				tmp.Weather.String = v
			case 9:
				tmp.Hunting.Valid = null
				if null {
					continue
				}
				tmp.Hunting.Duration, err = time.ParseDuration(v)
				if err != nil {
					return nil, err
				}
			case 10:
				tmp.Devotion.Valid = null
				if null {
					continue
				}
				tmp.Devotion.Duration, err = time.ParseDuration(v)
				if err != nil {
					return nil, err
				}
			case 11:
				tmp.Hobby.Valid = null
				if null {
					continue
				}
				tmp.Hobby.Duration, err = time.ParseDuration(v)
				if err != nil {
					return nil, err
				}
			case 12:
				tmp.WorkoutW.Valid = null
				if null {
					continue
				}
				tmp.WorkoutW.Duration, err = time.ParseDuration(v)
				if err != nil {
					return nil, err
				}
			case 13:
				tmp.WorkoutR.Valid = null
				if null {
					continue
				}
				tmp.WorkoutR.Float, err = strconv.ParseFloat(v, 64)
				if err != nil {
					return nil, err
				}
			case 14:
				tmp.WorkoutB.Valid = null
				if null {
					continue
				}
				tmp.WorkoutB.Int, err = strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
			}
		}
		res = append(res, tmp)
	}
	return res, nil
}
