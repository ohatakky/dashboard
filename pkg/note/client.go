package note

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// curl "https://note.com/api/v2/creators/ohatakky/contents?kind=note&page=1" | jq .

const (
	endpoint       = "https://note.com/api/v2"
	contentPathFmt = "creators/%s/contents?kind=note"
)

type Posts struct {
	Data struct {
		Contents []struct {
			ID                   int         `json:"id"`
			Type                 string      `json:"type"`
			Status               string      `json:"status"`
			Name                 string      `json:"name"`
			Description          interface{} `json:"description"`
			Price                int         `json:"price"`
			Key                  string      `json:"key"`
			Slug                 string      `json:"slug"`
			PublishAt            string      `json:"publishAt"`
			ThumbnailExternalURL string      `json:"thumbnailExternalUrl"`
			Eyecatch             string      `json:"eyecatch"`
			User                 struct {
				ID                     int         `json:"id"`
				Name                   string      `json:"name"`
				Urlname                string      `json:"urlname"`
				Nickname               string      `json:"nickname"`
				UserProfileImagePath   string      `json:"userProfileImagePath"`
				CustomDomain           interface{} `json:"customDomain"`
				DisableSupport         bool        `json:"disableSupport"`
				LikeAppealText         interface{} `json:"likeAppealText"`
				LikeAppealImage        interface{} `json:"likeAppealImage"`
				PurchaseAppealTextNote interface{} `json:"purchaseAppealTextNote"`
				TwitterNickname        string      `json:"twitterNickname"`
			} `json:"user"`
			CanRead        bool        `json:"canRead"`
			IsAuthor       bool        `json:"isAuthor"`
			ExternalURL    interface{} `json:"externalUrl"`
			CustomDomain   interface{} `json:"customDomain"`
			Body           string      `json:"body"`
			IsLimited      bool        `json:"isLimited"`
			IsTrial        bool        `json:"isTrial"`
			CanUpdate      bool        `json:"canUpdate"`
			TweetText      string      `json:"tweetText"`
			AdditionalAttr struct {
				Formatted bool `json:"formatted"`
			} `json:"additionalAttr"`
			IsRefund           bool          `json:"isRefund"`
			CommentCount       int           `json:"commentCount"`
			Likes              []interface{} `json:"likes"`
			LikeCount          int           `json:"likeCount"`
			AnonymousLikeCount int           `json:"anonymousLikeCount"`
			IsLiked            bool          `json:"isLiked"`
			DisableComment     bool          `json:"disableComment"`
			Hashtags           []struct {
				Hashtag struct {
					Name string `json:"name"`
				} `json:"hashtag"`
			} `json:"hashtags"`
			TwitterShareURL  string `json:"twitterShareUrl"`
			FacebookShareURL string `json:"facebookShareUrl"`
			Audio            struct {
			} `json:"audio"`
			Pictures               []interface{} `json:"pictures"`
			LimitedMessage         interface{}   `json:"limitedMessage"`
			Labels                 []interface{} `json:"labels"`
			PriorSale              interface{}   `json:"priorSale"`
			CanMultipleLimitedNote bool          `json:"canMultipleLimitedNote"`
			HasEmbeddedContent     bool          `json:"hasEmbeddedContent"`
			IsPinned               bool          `json:"isPinned"`
			PinnedUserNoteID       interface{}   `json:"pinnedUserNoteId"`
			IsTreasuredNote        bool          `json:"isTreasuredNote"`
			SpEyecatch             string        `json:"spEyecatch"`
			EnableBacktoDraft      bool          `json:"enableBacktoDraft"`
			NotificationMessages   []interface{} `json:"notificationMessages"`
			IsProfiled             bool          `json:"isProfiled"`
			IsForWork              bool          `json:"isForWork"`
			IsCircleDescription    bool          `json:"isCircleDescription"`
			NoteDraft              interface{}   `json:"noteDraft"`
		} `json:"contents"`
		IsLastPage bool `json:"isLastPage"`
		TotalCount int  `json:"totalCount"`
	} `json:"data"`
}

type Client struct {
	user string
}

func NewClient(user string) *Client {
	return &Client{
		user: user,
	}
}

func (c *Client) GetPosts() (*Posts, error) {
	u := endpoint + "/" + fmt.Sprintf(contentPathFmt, c.user)
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := Posts{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
