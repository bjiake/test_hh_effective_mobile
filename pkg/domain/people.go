package domain

type People struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	SurName    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
