package service

import (
	"context"
	"errors"
	"hh.ru/pkg/api/filter"
	"hh.ru/pkg/db"
	"hh.ru/pkg/domain"
	"log"
)

func (s service) GetCarFilter(ctx context.Context, filter *filter.Filter) ([]domain.Car, error) {
	carFilter, err := s.repoCar.Get(ctx, filter)
	if err != nil {
		return nil, err
	}
	return carFilter, nil
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

//func (s service) Get(ctx context.Context, id int64) (domain.Car, error) {
//	car, err := s.repoCar.GetByID(ctx, id)
//	if errors.Is(err, db.ErrNotExist) {
//		log.Printf("car: %d does not exist in the repository\n", id)
//		return domain.Car{}, err
//	} else if err != nil {
//		log.Println(err)
//		return domain.Car{}, err
//	}
//	return *car, nil
//}
