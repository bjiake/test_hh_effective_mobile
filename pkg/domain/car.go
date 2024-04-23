package domain

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

// Car represents a car entity
type Car struct {
	// ID is the unique identifier of the car
	ID int64 `json:"id" swagger:"description:Unique identifier of the car"`

	// RegNum is the registration number of the car
	RegNum string `json:"regNum" validate:"required,regNum" swagger:"description:Registration number of the car"`

	// Mark is the brand of the car
	Mark string `json:"mark" validate:"required" swagger:"description:Brand of the car"`

	// Model is the model of the car
	Model string `json:"model" validate:"required" swagger:"description:Model of the car"`

	// Year is the year of manufacture of the car
	Year int64 `json:"year" validate:"required,number" swagger:"description:Year of manufacture of the car"`

	// Owner is the ID of the car's owner
	Owner int64 `json:"owner" validate:"required" swagger:"description:ID of the car's owner"`
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

type UpdateCar struct {
	ID     int64  `json:"id"`
	RegNum string `json:"regNum" validate:"omitempty,regNum"`
	Mark   string `json:"mark" validate:"omitempty"`
	Model  string `json:"domain" validate:"omitempty"`
	Year   int64  `json:"year" validate:"omitempty"`
	Owner  int64  `json:"owner" validate:"omitempty"`
}

func (c UpdateCar) Validate() error {
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

// RequestCar represents a car request
type RequestCar struct {
	// ID is the unique identifier of the car
	ID int64 `json:"id" swagger:"description:Unique identifier of the car"`

	// RegNum is the registration number of the car
	RegNum string `json:"regNum" swagger:"description:Registration number of the car"`

	// Mark is the brand of the car
	Mark string `json:"mark" swagger:"description:Brand of the car"`

	// Model is the model of the car
	Model string `json:"model" swagger:"description:Model of the car"`

	// Year is the year of manufacture of the car
	Year int64 `json:"year" swagger:"description:Year of manufacture of the car"`

	// Owner is the owner of the car
	Owner People `json:"owner" swagger:"description:Owner of the car"`
}
