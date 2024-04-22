package interfaces

import (
	"context"
	"hh.ru/cmd/internal/model"
)

type RepositoryPeople interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, people model.People) (*model.People, error)
	All(ctx context.Context) ([]model.People, error)
	GetByID(ctx context.Context, id int64) (*model.People, error)
	Update(ctx context.Context, id int64, updatedPeo model.People) (*model.People, error)
	Delete(ctx context.Context, id int64) error
}
