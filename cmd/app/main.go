package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"hh.ru/pkg/config"
	"hh.ru/pkg/di"
	"log"
)

// @title hh.ru/test/mobile
// @version 1.0
// @description API Server for test work
// @host localhost:3000
// @BasePath /
func main() {

	cfg, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	server, diErr := di.InitializeAPI(cfg)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
