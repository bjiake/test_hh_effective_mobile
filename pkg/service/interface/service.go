package interfaces

import (
	"context"
	"hh.ru/pkg/api/filter"
	"hh.ru/pkg/domain"
)

type ServiceUseCase interface {
	// Car func
	GetCarByRegNum(ctx context.Context, regNum string) (*domain.RequestCar, error)
	GetCar(ctx context.Context, filter *filter.Car) ([]domain.RequestCar, error)
	CreateCar(ctx context.Context, car domain.Car) (*domain.RequestCar, error)
	DeleteCar(ctx context.Context, id int64) error
	UpdateCar(ctx context.Context, uCar domain.UpdateCar) (*domain.RequestCar, error)
	// People func
	CreatePeople(ctx context.Context, people domain.People) (*domain.People, error)
	DeletePeople(ctx context.Context, id int64) error
	GetPeople(ctx context.Context, filter *filter.People) ([]domain.People, error)
	UpdatePeople(ctx context.Context, uPeople domain.UpdatePeople) (*domain.People, error)
}
