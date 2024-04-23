package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"hh.ru/pkg/api/filter"
	"hh.ru/pkg/db"
	"hh.ru/pkg/domain"
	"log"
	"regexp"
)

func (s service) UpdateCar(ctx context.Context, uCar domain.UpdateCar) (*domain.RequestCar, error) {
	// find car object
	filterI := filter.Car{ID: &uCar.ID}
	temp, err := s.GetCar(ctx, &filterI)
	if err != nil {
		return nil, err
	}
	car := domain.Car{
		ID:     temp[0].ID,
		Model:  temp[0].Model,
		RegNum: temp[0].RegNum,
		Mark:   temp[0].Mark,
		Year:   temp[0].Year,
		Owner:  temp[0].Owner.ID,
	}

	if uCar.RegNum != "" && uCar.RegNum != car.RegNum {
		car.RegNum = uCar.RegNum
	}
	if uCar.Mark != "" && uCar.Mark != car.Mark {
		car.Mark = uCar.Mark
	}
	if uCar.Model != "" && uCar.Model != car.Model {
		car.Model = uCar.Model
	}
	if uCar.Year != 0 && uCar.Year != car.Year {
		car.Year = uCar.Year
	}
	if uCar.Owner != 0 && uCar.Owner != car.Owner {
		car.Owner = uCar.Owner
	}

	if err := car.Validate(); err != nil {
		return nil, err
	}

	requestCar, err := s.repoCar.Update(ctx, car.ID, car)
	if errors.Is(err, db.ErrDuplicate) {
		return nil, fmt.Errorf("record: %+v already exists\n", uCar)
	} else if errors.Is(err, db.ErrUpdateFailed) {
		return nil, fmt.Errorf("update of record: %+v failed", uCar)
	} else if err != nil {
		return nil, err
	}

	return requestCar, err
}

// Custom validation function for Latin and Cyrillic characters
func validateLatinCyrillic(fl validator.FieldLevel) bool {
	// Regular expression to match Latin and Cyrillic characters
	regex := regexp.MustCompile(`^[\p{Latin}\p{Cyrillic}\s]+$`)
	return regex.MatchString(fl.Field().String())
}

func (s service) UpdatePeople(ctx context.Context, uPeople domain.UpdatePeople) (*domain.People, error) {
	filterI := filter.People{ID: &uPeople.ID}
	todo, err := s.GetPeople(ctx, &filterI)
	if err != nil {
		return nil, err
	}

	if err = uPeople.Validate(); err != nil {
		return nil, err
	}

	if uPeople.Name != "" {
		todo[0].Name = uPeople.Name
	}
	if uPeople.SurName != "" {
		todo[0].SurName = uPeople.SurName
	}
	if uPeople.Patronymic != "" {
		todo[0].Patronymic = uPeople.Patronymic
	}

	resultPeople, err := s.repoPeople.Update(ctx, todo[0].ID, todo[0])
	if errors.Is(err, db.ErrDuplicate) {
		return resultPeople, fmt.Errorf("record: %+v already exists\n", uPeople)
	} else if errors.Is(err, db.ErrUpdateFailed) {
		return resultPeople, fmt.Errorf("update of record: %+v failed", uPeople)
	} else if err != nil {
		log.Println(err)
		return resultPeople, err
	}

	return resultPeople, err
}
