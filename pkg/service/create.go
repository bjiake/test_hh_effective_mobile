package service

import (
	"context"
	"errors"
	"fmt"
	"hh.ru/pkg/db"
	"hh.ru/pkg/domain"
	"log"
)

func (s service) CreateCar(ctx context.Context, car domain.Car) (*domain.Car, error) {
	if err := car.Validate(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	resultEntity, err := s.repoCar.Create(ctx, car)
	if errors.Is(err, db.ErrDuplicate) {
		return nil, fmt.Errorf("record: %+v already exists\n", car)
	} else if err != nil {
		log.Println(err)
		return nil, err
	}

	return resultEntity, err
}

func (s service) CreatePeople(ctx context.Context, people domain.People) (*domain.People, error) {
	if err := people.Validate(); err != nil {
		return nil, err
	}

	resultEntity, err := s.repoPeople.Create(ctx, people)
	if errors.Is(err, db.ErrDuplicate) {
		return nil, fmt.Errorf("record: %+v already exists\n", people)
	} else if err != nil {
		return nil, err
	}

	return resultEntity, err
}
