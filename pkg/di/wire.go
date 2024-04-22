//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "hh.ru/pkg/api"
	"hh.ru/pkg/api/handler"
	"hh.ru/pkg/config"
	"hh.ru/pkg/db"
	"hh.ru/pkg/service"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectToBD, service.NewService, handler.NewHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
