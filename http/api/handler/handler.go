package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"hh.ru/cmd/internal/model"
	services "hh.ru/cmd/internal/service"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) FindCarByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.ParseInt(paramsId, 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	car, err := h.service.GetCar(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, car)
}

func (h *Handler) CreateCar(c *gin.Context) {
	var car model.Car

	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error create": err.Error()})
		log.Println("error binding car:", err.Error())
		return
	}

	car, err := h.service.Create(c, car)
}
