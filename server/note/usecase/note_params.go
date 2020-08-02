package usecase

type PostsDailyResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}
