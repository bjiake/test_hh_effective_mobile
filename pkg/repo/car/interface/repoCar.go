package interfaces

import (
	"context"
	"hh.ru/pkg/domain"
)

type CarRepository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, car domain.Car) (*domain.Car, error)
	All(ctx context.Context) ([]domain.Car, error)
	GetByID(ctx context.Context, id int64) (*domain.Car, error)
	Update(ctx context.Context, id int64, updatedCar domain.Car) (*domain.Car, error)
	Delete(ctx context.Context, id int64) error
}
