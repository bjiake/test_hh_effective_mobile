package service

import (
	"context"
	"errors"
	"fmt"
	"hh.ru/pkg/db"
)

func (s service) Delete(ctx context.Context, id int64) error {
	err := s.repoCar.Delete(ctx, id)
	if errors.Is(err, db.ErrDeleteFailed) {
		fmt.Printf("delete of record: %d failed", id)
		return err
	} else if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
