package service

import (
	"context"
	"errors"
	"hh.ru/pkg/db"
	"hh.ru/pkg/domain"
	"log"
)

func (s service) GetCar(ctx context.Context, id int64) (domain.Car, error) {
	car, err := s.repoCar.GetByID(ctx, id)
	if errors.Is(err, db.ErrNotExist) {
		log.Printf("car: %d does not exist in the repository\n", id)
		return domain.Car{}, err
	} else if err != nil {
		log.Println(err)
		return domain.Car{}, err
	}
	return *car, nil
}

func (s service) GetPeople(ctx context.Context, id int64) (domain.People, error) {
	people, err := s.repoPeople.GetByID(ctx, id)
	if errors.Is(err, db.ErrNotExist) {
		log.Printf("people: %d does not exist in the repository\n", id)
		return domain.People{}, err
	} else if err != nil {
		log.Println(err)
		return domain.People{}, err
	}
	return *people, nil
}
