package service

import (
	"errors"
	"github.com/fir1/rest-api/pkg/erru"
	"github.com/gin-gonic/gin"
	"hh.ru/cmd/internal/repo"
	"hh.ru/cmd/internal/repo/car"
	"hh.ru/cmd/internal/repo/people"
	"log"
	"github.com/asaskevich/govalidator"
)



func GetInfo(context *gin.Context) {
	regNum, err := parseParam(context)
	if err != nil {
		context.JSON(400, err.Error())
		return
	}

	result, err := repo.
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
