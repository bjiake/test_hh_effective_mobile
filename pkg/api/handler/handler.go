package handler

import (
	gin "github.com/gin-gonic/gin"
	"hh.ru/pkg/api/filter"
	"hh.ru/pkg/domain"
	services "hh.ru/pkg/service/interface"
	"log"
	"net/http"
	"regexp"
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

// GetCarByRegNum returns a car by registration number
// @Summary Get a car by registration number
// @Description Get a car by its registration number
// @Tags Car
// @Accept  json
// @Produce  json
// @Param   regNum path     string true "Registration number of the car"
// @Success 200 {object} domain.Car "OK"
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /car/{regNum} [get]
func (h *Handler) GetCarByRegNum(c *gin.Context) {
	paramsRegNum := c.Param("regNum")
	if paramsRegNum == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "params regNum can't be empty"})
		log.Println("params regNum can't be empty")
		return
	}

	_, err := regexp.MatchString(`^[A-ZА-Я]\d{3}[A-ZА-Я]{2}\d{3}$`, paramsRegNum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid params regNum"})
		log.Println("invalid params regNum")
		return
	}

	car, err := h.service.GetCarByRegNum(c.Request.Context(), paramsRegNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, car)
}

// GetCar godoc
// @Summary Get cars
// @Description get cars by filter and with pagination
// @Tags Car
// @Accept  json
// @Produce  json
// @Param   limit   query    int     false  "Limit"   minimum(0)
// @Param   offset  query    int     false  "Offset"  minimum(1)
// @Param   id      query    int     false  "ID of the car"
// @Param   regNum  query    string  false  "Registration number of the car"
// @Param   mark    query    string  false  "Brand of the car"
// @Param   model   query    string  false  "Model of the car"
// @Param   year    query    int     false  "Year of the car"
// @Param   owner   query    int     false  "Owner ID of the car"
// @Success 200 {object} gin.H "Successful retrieval of cars"
// @Failure 400 {object} gin.H "Error: Bad Request"
// @Failure 500 {object} gin.H "Error: Internal Server Error"
// @Router /cars [get]
func (h *Handler) GetCar(c *gin.Context) {
	var filterI filter.Car
	if err := c.BindQuery(&filterI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error bind Get filter": err.Error()})
		return
	}
	var paginationI filter.Pagination
	if err := c.BindQuery(&paginationI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error bind Get pagination": err.Error()})
		return
	}
	cars, err := h.service.GetCar(c.Request.Context(), &filterI, &paginationI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// CreateCar godoc
// @Summary Create a new car
// @Description Create a new car with the provided data
// @Tags Car
// @Accept json
// @Produce json
// @Param car body domain.Car true "Car data"
// @Success 201 {object} domain.Car "Created"
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /car [post]
func (h *Handler) CreateCar(c *gin.Context) {
	var createCar domain.Car

	if err := c.BindJSON(&createCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error create": err.Error()})
		log.Println("error binding car:", err.Error())
		return
	}

	car, err := h.service.CreateCar(c.Request.Context(), createCar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println("error create car:", err.Error())
		return
	}
	c.JSON(http.StatusCreated, car)
}

// UpdateCar godoc
// @Summary Update a car
// @Description Update an existing car with the provided data
// @Tags Car
// @Accept json
// @Produce json
// @Param car body domain.UpdateCar true "Car update data"
// @Success 200 {object} domain.Car "OK"
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /car [put]
func (h *Handler) UpdateCar(c *gin.Context) {
	var updateCar domain.UpdateCar
	if err := c.BindJSON(&updateCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error update": err.Error()})
		log.Println("error binding car:", err.Error())
		return
	}
	car, err := h.service.UpdateCar(c.Request.Context(), updateCar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println("error update car:", err.Error())
		return
	}
	c.JSON(http.StatusOK, car)
}

// DeleteCar godoc
// @Summary Delete a car
// @Description Delete a car by ID
// @Tags Car
// @Accept json
// @Produce plain
// @Param id path integer true "Car ID"
// @Success 204 "No Content"
// @Failure 500 {object} string "Internal Server Error"
// @Router /car/{id} [delete]
func (h *Handler) DeleteCar(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.ParseInt(paramsId, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}

	err = h.service.DeleteCar(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"deleted": true})
}

// GetPeople godoc
// @Summary Get people
// @Description Get a list of people with optional filtering
// @Tags People
// @Accept json
// @Produce json
// @Param id query integer false "Filter by ID"
// @Param name query string false "Filter by name"
// @Param surName query string false "Filter by surname"
// @Param patronymic query string false "Filter by patronymic"
// @Success 200 {array} domain.People "OK"
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /people [get]
func (h *Handler) GetPeople(c *gin.Context) {
	var filterI filter.People
	if err := c.BindQuery(&filterI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err.Error())
		return
	}
	cars, err := h.service.GetPeople(c.Request.Context(), &filterI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// CreatePeople godoc
// @Summary Create a new person
// @Description Create a new person with the provided data
// @Tags People
// @Accept json
// @Produce json
// @Param people body domain.People true "Person data"
// @Success 201 {object} domain.People "Created"
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /people [post]
func (h *Handler) CreatePeople(c *gin.Context) {
	var createCar domain.People

	if err := c.BindJSON(&createCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error create": err.Error()})
		log.Println(err.Error())
		return
	}

	car, err := h.service.CreatePeople(c.Request.Context(), createCar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusCreated, car)
}

// UpdatePeople godoc
// @Summary Update a person
// @Description Update an existing person with the provided data
// @Tags People
// @Accept json
// @Produce json
// @Param people body domain.UpdatePeople true "Person update data"
// @Success 200 {object} domain.People "OK"
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /people [put]
func (h *Handler) UpdatePeople(c *gin.Context) {
	var updatePeople domain.UpdatePeople
	if err := c.BindJSON(&updatePeople); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error update": err.Error()})
		log.Println(err.Error())
		return
	}
	car, err := h.service.UpdatePeople(c.Request.Context(), updatePeople)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, car)
}

// DeletePeople godoc
// @Summary Delete a person
// @Description Delete a person by ID
// @Tags People
// @Accept json
// @Produce plain
// @Param id path integer true "Person ID"
// @Success 204 "No Content"
// @Failure 500 {object} string "Internal Server Error"
// @Router /people/{id} [delete]
func (h *Handler) DeletePeople(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.ParseInt(paramsId, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}

	err = h.service.DeletePeople(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
