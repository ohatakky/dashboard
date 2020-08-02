package usecase

type SubmissionsDailyResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}
