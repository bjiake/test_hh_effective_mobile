package service

import (
	"context"
	"errors"
	"fmt"
	"hh.ru/pkg/db"
	"hh.ru/pkg/domain"
	"log"
)

func (s service) Create(ctx context.Context, car domain.Car) (*domain.Car, error) {
	if err := car.Validate(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	resultEntity, err := s.repoCar.Create(ctx, car)
	if errors.Is(err, db.ErrDuplicate) {
		fmt.Printf("record: %+v already exists\n", car)
		return nil, err
	} else if err != nil {
		log.Println(err)
		return nil, err
	}

	return resultEntity, err
}
