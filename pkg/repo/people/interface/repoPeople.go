package interfaces

import (
	"context"
	"hh.ru/pkg/api/filter"
	"hh.ru/pkg/domain"
)

type PeopleRepository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, people domain.People) (*domain.People, error)
	Get(ctx context.Context, filter *filter.People) ([]domain.People, error)
	Update(ctx context.Context, id int64, updatedPeo domain.People) (*domain.People, error)
	Delete(ctx context.Context, id int64) error
}
