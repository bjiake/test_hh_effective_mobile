package service

import (
	"context"
)

func (s Service) Delete(ctx context.Context, id int64) error {
	err := s.repoCar.Delete(ctx, id)
	if err != nil {
		return err
	}

	return err
}
