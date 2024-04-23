package service

import (
	"context"
	"hh.ru/pkg/api/filter"
	"hh.ru/pkg/domain"
)

func (s service) GetCar(ctx context.Context, filter *filter.Car) ([]domain.Car, error) {
	cars, err := s.repoCar.Get(ctx, filter)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (s service) GetPeople(ctx context.Context, filter *filter.People) ([]domain.People, error) {
	peoples, err := s.repoPeople.Get(ctx, filter)
	if err != nil {
		return nil, err
	}
	return peoples, nil
}
