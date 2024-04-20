package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hh.ru/cmd/internal/repo"
	"log"
)

func GetInfo(context *gin.Context) {
	regNum, err := parseParam(context)
	if err != nil {
		context.JSON(400, err.Error())
		return
	}

	result, err := repo.NewPostSQLClassicRepository(db)
}

func parseParam(context *gin.Context) (string, error) {
	regNum := context.Param("regNum")
	if regNum == "" {
		message := "invalid parse regNum"
		log.Println(message)
		return "", errors.New(message)
	}
	return regNum, nil
}
