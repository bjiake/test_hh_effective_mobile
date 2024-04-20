package model

type Car struct {
	ID     int64  `json:"id"`
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int32  `json:"year"`
	Owner  int64  `json:"owner"`
}
