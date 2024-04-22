package service

import (
	"hh.ru/cmd/internal/repo/car"
	"hh.ru/cmd/internal/repo/people"
)

type Service struct {
	repoCar    car.RepositoryCar
	repoPeople people.RepositoryPeople
}

func NewService(c car.RepositoryCar, p people.RepositoryPeople) Service {
	return Service{
		repoCar:    c,
		repoPeople: p,
	}
}
