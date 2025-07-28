package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"

	"github.com/gin-gonic/gin"
)

type CarModelHandler struct {
	service *services.CarModelService
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	return &CarModelHandler{
		service: services.NewCarModelService(cfg),
	}
}

// CreateCarModel godoc
// @Summary Create a CarModel
// @Description Create a CarModel
// @Tags CarModels
// @Accept json
// @Produce json
// @Param Request body dto.CreateCarModelRequest true "Create a CarModel"
// @Success 201 {object} dto.CarModelResponse "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-models/ [post]
// @Security AuthBearer
func (h *CarModelHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

//Update

// UpdateCarModel godoc
// @Summary Update a CarModel
// @Description Update a CarModel
// @Tags CarModels
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelRequest true "Update a CarModel"
// @Success 200 {object} dto.CarModelResponse "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-models/{id} [put]
// @Security AuthBearer
func (h *CarModelHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

//Deleted

// DeleteCarModel godoc
// @Summary Delete a CarModel
// @Description Delete a CarModel
// @Tags CarModels
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Deleted successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-models/{id} [delete]
// @Security AuthBearer
func (h *CarModelHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

//get

// GetCarModel godoc
// @Summary Get a CarModel
// @Description Get a CarModel
// @Tags CarModels
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} dto.CarModelResponse "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-models/{id} [get]
// @Security AuthBearer
func (h *CarModelHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCarModels godoc
// @Summary Filter Countries by pagination & criteria
// @Description Returns paginated list of countries based on filters
// @Tags CarModels
// @Accept json
// @Produce json
// @Param Request body  dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-models/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
