package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hh.ru/cmd/internal/model"
	"hh.ru/cmd/internal/repo/car"
	"hh.ru/cmd/internal/repo/people"
	"hh.ru/pkg/erru"
	"regexp"
	"strings"
)

type Service struct {
	repoCar    car.RepositoryCar
	repoPeople people.RepositoryPeople
}

func NewService(c car.RepositoryCar, p people.RepositoryPeople) Service {
	return Service{
		repoCar:    c,
		repoPeople: p,
	}
}

type CreateParamsCar struct {
	RegNum string `json:"regNum" validate:"required,regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int32  `json:"year" validate:"gte=1600,lte=2024"`
	Owner  int64  `json:"owner"`
}

func (c CreateParamsCar) Validate() error {
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

func (s Service) Create(ctx *gin.Context, params CreateParamsCar) (*model.Car, error) {
	if err := params.Validate(); err != nil {
		return nil, erru.ErrArgument{Wrapped: err}
	}

	entity := model.Car{
		RegNum: params.RegNum,
		Mark:   params.Mark,
		Model:  params.Model,
		Year:   params.Year,
		Owner:  params.Owner,
	}

	resultEntity, err := s.repoCar.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	return resultEntity, err
}
