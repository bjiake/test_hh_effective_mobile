package service

import (
	"context"
	"hh.ru/pkg/api/filter"
	"hh.ru/pkg/domain"
)

func (s service) GetCarByRegNum(ctx context.Context, regNum string) (*domain.RequestCar, error) {
	car, err := s.repoCar.GetByRegNum(ctx, regNum)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (s service) GetCar(ctx context.Context, filterI *filter.Car, pagination *filter.Pagination) ([]domain.RequestCar, error) {
	cars, err := s.repoCar.Get(ctx, filterI, pagination)
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
