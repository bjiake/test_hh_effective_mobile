package res

import (
	"github.com/fir1/rest-api/configs"
	"github.com/fir1/rest-api/http/rest/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"hh.ru/pkg/db"
)

type Server struct {
	logger *logrus.Logger
	router *mux.Router
	config configs.Config
}

func NewServer() (*Server, error) {
	cnf, err := configs.NewParsedConfig()
	if err != nil {
		return nil, err
	}

	database, err := db.ConnectToBD()
	if err != nil {
		return nil, err
	}

	log := NewLogger()
	router := mux.NewRouter()
	handlers.Register(router, log, database)

	s := Server{
		logger: log,
		config: cnf,
		router: router,
	}
	return &s, nil
}
