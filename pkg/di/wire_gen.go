package di

import (
	"hh.ru/cmd/internal/repo/car"
	"hh.ru/cmd/internal/repo/people"
	"hh.ru/cmd/internal/service"
	"hh.ru/pkg/config"
	"hh.ru/pkg/db"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	bd, err := db.ConnectToBD(cfg)
	if err != nil {
		return nil, err
	}
	userService := service.NewService(car.NewPgSqlCarRepository(bd), people.NewPgSqlPeopleRepository(bd))
	userUseCase := usecase.NewUserUseCase(userService)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}
