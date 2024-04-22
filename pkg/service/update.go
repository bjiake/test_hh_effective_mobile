package service

import (
	"context"
	"errors"
	"fmt"
	"hh.ru/pkg/db"
	"hh.ru/pkg/domain"
	"log"
)

func (s service) Update(ctx context.Context, updateCar domain.Car) (*domain.Car, error) {
	// find todo object
	todo, err := s.GetCar(ctx, updateCar.ID)
	if err != nil {
		return nil, err
	}

	if updateCar.RegNum != "" && updateCar.Validate() == nil {
		todo.RegNum = updateCar.RegNum
	}
	if updateCar.Mark != "" {
		todo.Mark = updateCar.Mark
	}
	if updateCar.Model != "" {
		todo.Model = updateCar.Model
	}
	if updateCar.Year != 0 {
		todo.Year = updateCar.Year
	}
	if updateCar.Owner != 0 {
		todo.Owner = updateCar.Owner
	}

	resultCar, err := s.repoCar.Update(ctx, todo.ID, todo)
	if errors.Is(err, db.ErrDuplicate) {
		fmt.Printf("record: %+v already exists\n", updateCar)
		return resultCar, err
	} else if errors.Is(err, db.ErrUpdateFailed) {
		fmt.Printf("update of record: %+v failed", updateCar)
		return resultCar, err
	} else if err != nil {
		log.Println(err)
		return resultCar, err
	}

	return resultCar, err
}
