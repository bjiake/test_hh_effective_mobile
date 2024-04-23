package handler

import (
	"github.com/gin-gonic/gin"
	"hh.ru/pkg/api/filter"
	"hh.ru/pkg/domain"
	services "hh.ru/pkg/service/interface"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	service services.ServiceUseCase
}

func NewHandler(service services.ServiceUseCase) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) FindCarFilter(c *gin.Context) {
	var filterI filter.Filter
	if err := c.BindQuery(&filterI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cars, err := h.service.GetCarFilter(c.Request.Context(), &filterI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cars})
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

	car, err := h.service.GetCar(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, car)
}

func (h *Handler) CreateCar(c *gin.Context) {
	var createCar domain.Car

	if err := c.BindJSON(&createCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error create": err.Error()})
		log.Println("error binding car:", err.Error())
		return
	}

	car, err := h.service.Create(c.Request.Context(), createCar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println("error create car:", err.Error())
		return
	}
	c.JSON(http.StatusCreated, car)
}
func (h *Handler) UpdateCar(c *gin.Context) {
	var updateCar domain.Car
	if err := c.BindJSON(&updateCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error update": err.Error()})
		log.Println("error binding car:", err.Error())
		return
	}
	car, err := h.service.Update(c.Request.Context(), updateCar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println("error update car:", err.Error())
		return
	}
	c.JSON(http.StatusOK, car)
}

func (h *Handler) DeleteCar(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.ParseInt(paramsId, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println("error param id:", err.Error())
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println("error delete car:", err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
