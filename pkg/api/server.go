package api

import (
	"github.com/gin-gonic/gin"
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
	//engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	engine.GET("car/:id", userHandler.FindCarByID)
	engine.PUT("car", userHandler.UpdateCar)
	engine.POST("car", userHandler.CreateCar)
	engine.DELETE("car/:id", userHandler.DeleteCar)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
