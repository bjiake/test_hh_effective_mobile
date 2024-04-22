package service

import (
	"context"
	"hh.ru/cmd/internal/model"
)

func (s Service) GetCar(ctx context.Context, id int64) (model.Car, error) {
	car, err := s.repoCar.GetByID(ctx, id)
	if err != nil {
		return model.Car{}, err
	}
	return *car, nil
}

func (s Service) GetPeople(ctx context.Context, id int64) (model.People, error) {
	people, err := s.repoPeople.GetByID(ctx, id)
	if err != nil {
		return model.People{}, err
	}
	return *people, nil
}
