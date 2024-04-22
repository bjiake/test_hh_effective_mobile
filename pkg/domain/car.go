package domain

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

/*
 Регистрация кастомного тега
validate := validator.New()
validate.RegisterValidation("regNum", validateRegNum)
*/

type Car struct {
	ID     int64  `json:"id"`
	RegNum string `json:"regNum" validate:"required,regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"domain"`
	Year   int32  `json:"year" validate:"gte=1900,lte=2024"`
	Owner  int64  `json:"owner"`
}

func (c Car) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("regNum", validateRegNum)
	if err != nil {
		return err
	}

	err = validate.Struct(c)
	if err != nil {
		// Обработка ошибок валидации
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return fmt.Errorf("car validation errors: %s", strings.Join(validationErrors, ", "))
	}
	return err
}

func validateRegNum(fl validator.FieldLevel) bool {
	regNum := fl.Field().String()
	// Регулярное выражение для латинских и кириллических букв
	matched, err := regexp.MatchString(`^[A-ZА-Я]\d{3}[A-ZА-Я]{2}\d{3}$`, regNum)
	if err != nil {
		return false
	}
	return matched
}
