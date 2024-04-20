package model

type People struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
