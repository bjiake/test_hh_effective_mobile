package model

/*
 Регистрация кастомного тега
validate := validator.New()
validate.RegisterValidation("regNum", validateRegNum)
*/

type Car struct {
	ID     int64  `json:"id"`
	RegNum string `json:"regNum" validate:"required,regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int32  `json:"year" validate:"gte=1900,lte=2024"`
	Owner  int64  `json:"owner"`
}
