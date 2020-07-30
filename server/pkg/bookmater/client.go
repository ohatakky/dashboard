package bookmater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// curl 'https://bookmeter.com/users/1064403/reviews.json?offset=0&limit=100' | jq .

const (
	endpoint   = "https://bookmeter.com"
	reviewsFmt = "users/%s/reviews.json?offset=0&limit=100"
)

type Reviews struct {
	Metadata struct {
		Sort   string `json:"sort"`
		Order  string `json:"order"`
		Offset int    `json:"offset"`
		Limit  int    `json:"limit"`
		Count  int    `json:"count"`
	} `json:"metadata"`
	Resources []struct {
		ID         int    `json:"id"`
		Path       string `json:"path"`
		Deletable  bool   `json:"deletable"`
		ContentTag string `json:"content_tag"`
		Content    string `json:"content"`
		CreatedAt  string `json:"created_at"`
		Highlight  bool   `json:"highlight"`
		Newly      bool   `json:"newly"`
		Contents   struct {
			ImageURL interface{} `json:"image_url"`
			Book     struct {
				ID         int    `json:"id"`
				Path       string `json:"path"`
				AmazonUrls struct {
					Outline      string `json:"outline"`
					Registration string `json:"registration"`
					WishBook     string `json:"wish_book"`
				} `json:"amazon_urls"`
				Title             string `json:"title"`
				ImageURL          string `json:"image_url"`
				RegistrationCount int    `json:"registration_count"`
				Page              int    `json:"page"`
				Original          bool   `json:"original"`
				IsAdvertisable    bool   `json:"is_advertisable"`
				Author            struct {
					Name string `json:"name"`
					Path string `json:"path"`
				} `json:"author"`
			} `json:"book"`
		} `json:"contents"`
		User struct {
			ID    int    `json:"id"`
			Path  string `json:"path"`
			Name  string `json:"name"`
			Image string `json:"image"`
		} `json:"user"`
		Nice struct {
			Path   string `json:"path"`
			Count  int    `json:"count"`
			Marked bool   `json:"marked"`
		} `json:"nice"`
		Netabare struct {
			Netabare       bool `json:"netabare"`
			DisplayContent bool `json:"display_content"`
			DisplayComment bool `json:"display_comment"`
			IsClicked      bool `json:"is_clicked"`
		} `json:"netabare"`
		NetabareDisplaySetting struct {
			ShouldDisplayIcon    bool `json:"should_display_icon"`
			ShouldGrayOutReview  bool `json:"should_gray_out_review"`
			ShouldDisplayComment bool `json:"should_display_comment"`
		} `json:"netabare_display_setting"`
		Comments struct {
			Path     string `json:"path"`
			Metadata struct {
				Sort   string `json:"sort"`
				Order  string `json:"order"`
				Offset int    `json:"offset"`
				Limit  int    `json:"limit"`
				Count  int    `json:"count"`
			} `json:"metadata"`
			Resources []interface{} `json:"resources"`
		} `json:"comments"`
	} `json:"resources"`
}

type Client struct {
	user string
}

func NewClient(user string) *Client {
	return &Client{
		user: user,
	}
}

func (c *Client) GetReviews() (*Reviews, error) {
	u := endpoint + "/" + fmt.Sprintf(reviewsFmt, c.user)
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := Reviews{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
