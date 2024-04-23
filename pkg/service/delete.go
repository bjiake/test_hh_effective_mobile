package service

import (
	"context"
	"errors"
	"fmt"
	"hh.ru/pkg/db"
	"log"
)

func (s service) DeleteCar(ctx context.Context, id int64) error {
	err := s.repoCar.Delete(ctx, id)
	if errors.Is(err, db.ErrDeleteFailed) {
		return fmt.Errorf("delete of record: %d failed", id)
	} else if err != nil {
		return err
	}
	return nil
}

func (s service) DeletePeople(ctx context.Context, id int64) error {
	err := s.repoPeople.Delete(ctx, id)
	if errors.Is(err, db.ErrDeleteFailed) {
		return fmt.Errorf("delete of record: %d failed", id)
	} else if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
