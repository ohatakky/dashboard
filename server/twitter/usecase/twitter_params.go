package usecase

type TweetsDailyResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}
