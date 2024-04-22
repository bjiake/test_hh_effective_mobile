package service

import (
	interfacesCar "hh.ru/pkg/repo/car/interface"
	interfacesPeople "hh.ru/pkg/repo/people/interface"
	interfaces "hh.ru/pkg/service/interface"
)

type service struct {
	repoCar    interfacesCar.CarRepository
	repoPeople interfacesPeople.PeopleRepository
}

func NewService(c interfacesCar.CarRepository, p interfacesPeople.PeopleRepository) interfaces.ServiceUseCase {
	return service{
		repoCar:    c,
		repoPeople: p,
	}
}
