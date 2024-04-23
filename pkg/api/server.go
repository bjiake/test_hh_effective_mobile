package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "hh.ru/docs"
	"hh.ru/pkg/api/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.Handler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	// to build swagger for me :) C:\Users\suyd\go\bin\swag init  --parseDependency -g ./cmd/app/main.go
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//Car methods
	engine.GET("info/:regNum", userHandler.GetCarByRegNum)
	engine.GET("car", userHandler.GetCar)
	engine.PUT("car", userHandler.UpdateCar)
	engine.POST("car", userHandler.CreateCar)
	engine.DELETE("car/:id", userHandler.DeleteCar)

	//People methods
	engine.GET("people", userHandler.GetPeople)
	engine.PUT("people", userHandler.UpdatePeople)
	engine.POST("people", userHandler.CreatePeople)
	engine.DELETE("people/:id", userHandler.DeletePeople)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
