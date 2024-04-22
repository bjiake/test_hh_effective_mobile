package service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"hh.ru/cmd/internal/model"
	"strings"
)

type UpdateParamsCar struct {
	ID     int64  `json:"id" valid:"Required"`
	RegNum string `json:"regNum" validate:"required,regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int32  `json:"year" validate:"gte=1600,lte=2024"`
	Owner  int64  `json:"owner"`
}

func (c UpdateParamsCar) Validate() error {
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

func (s Service) Update(ctx context.Context, params UpdateParamsCar) (*model.Car, error) {
	// find todo object
	todo, err := s.GetCar(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	if params.RegNum != "" && params.Validate() == nil {
		todo.RegNum = params.RegNum
	}
	if params.Mark != "" {
		todo.Mark = params.Mark
	}
	if params.Model != "" {
		todo.Model = params.Model
	}
	if params.Year != 0 {
		todo.Year = params.Year
	}
	if params.Owner != 0 {
		todo.Owner = params.Owner
	}

	resultCar, err := s.repoCar.Update(ctx, todo.ID, todo)
	if err != nil {
		return nil, err
	}

	return resultCar, err
}
