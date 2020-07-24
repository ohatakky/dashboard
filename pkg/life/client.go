package life

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ohatakky/dashboard/project/configs"
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

// todo: サーバー起動時にglobal変数にパースしておく。GETAPIは参照するだけ
// var

type Record struct {
	Date        time.Time     // 日付
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

func (c *Client) download() error {
	uFmt := "https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s"
	u := fmt.Sprintf(uFmt, configs.E.Life.SpreadsheetID, configs.E.Life.SheetName)
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	f, err := os.OpenFile("./life.csv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Records() ([]Record, error) {
	err := c.download()
	if err != nil {
		return nil, err
	}

	ff, err := os.OpenFile("./life.csv", os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer ff.Close()
	r := csv.NewReader(ff)
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

		// mapping to struct
		tmp := Record{}
		for i, v := range record {
			if v == "" {
				continue
			}
			switch i {
			case 0:
				tmp.Date, err = time.Parse(dateLayout, v)
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