package router

import (
	"github.com/gin-gonic/gin"
	"hh.ru/cmd/internal/service"
	"log"
)

func App() {
	router := gin.Default()

	router.GET("/info/:regNum", service.GetInfo)

	err := router.Run(":8088")
	if err != nil {
		log.Fatal(err)
		return
	}
}
