package interfaces

import (
	"context"
	"hh.ru/pkg/domain"
)

type ServiceUseCase interface {
	Create(ctx context.Context, car domain.Car) (*domain.Car, error)
	Delete(ctx context.Context, id int64) error
	GetCar(ctx context.Context, id int64) (domain.Car, error)
	Update(ctx context.Context, updateCar domain.Car) (*domain.Car, error)
}
