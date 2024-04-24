package interfaces

import (
	"context"
	"hh.ru/pkg/api/filter"
	"hh.ru/pkg/domain"
)

type CarRepository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, car domain.Car) (*domain.RequestCar, error)
	GetByRegNum(ctx context.Context, regNum string) (*domain.RequestCar, error)
	Get(ctx context.Context, filter *filter.Car, pagination *filter.Pagination) ([]domain.RequestCar, error)
	Update(ctx context.Context, id int64, updatedCar domain.Car) (*domain.RequestCar, error)
	Delete(ctx context.Context, id int64) error
}
