package db

import (
	"database/sql"
	"fmt"
	"hh.ru/pkg/config"
)

// ConnectToBD Подключение к PostgresSql по app.env
func ConnectToBD(cfg config.Config) (*sql.DB, error) {
	// Формирование строки подключения из конфига
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.PsqlUser, cfg.PsqlPass, cfg.PsqlHost, cfg.PsqlPort, cfg.PsqlDBName)

	// Подключение к БД
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
