package di

import (
	"hh.ru/cmd/internal/service"
	"hh.ru/pkg/db"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectToBD, service.NewService, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}