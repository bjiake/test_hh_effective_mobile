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

func (s service) UpdateCar(ctx context.Context, uCar domain.UpdateCar) (*domain.Car, error) {
	// find todo object
	filterI := filter.Car{ID: &uCar.ID}
	todo, err := s.GetCar(ctx, &filterI)
	if err != nil {
		return nil, err
	}

	if err := uCar.Validate(); err != nil {
		return nil, err
	}

	if uCar.RegNum != "" && uCar.Validate() == nil {
		todo[0].RegNum = uCar.RegNum
	}
	if uCar.Mark != "" {
		todo[0].Mark = uCar.Mark
	}
	if uCar.Model != "" {
		todo[0].Model = uCar.Model
	}
	if uCar.Year != 0 {
		todo[0].Year = uCar.Year
	}
	if uCar.Owner != 0 {
		todo[0].Owner = uCar.Owner
	}

	resultCar, err := s.repoCar.Update(ctx, todo[0].ID, todo[0])
	if errors.Is(err, db.ErrDuplicate) {
		return resultCar, fmt.Errorf("record: %+v already exists\n", uCar)
	} else if errors.Is(err, db.ErrUpdateFailed) {
		return resultCar, fmt.Errorf("update of record: %+v failed", uCar)
	} else if err != nil {
		return resultCar, err
	}

	return resultCar, err
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
