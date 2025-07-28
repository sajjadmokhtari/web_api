package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"

	"github.com/gin-gonic/gin"
)

type CarTypeHandler struct {
	service *services.CarTypeService
}

func NewCarTypeHandler(cfg *config.Config) *CarTypeHandler {
	return &CarTypeHandler{
		service: services.NewCarTypeService(cfg),
	}
}

// CreateCarType godoc
// @Summary Create a CarType
// @Description Create a CarType
// @Tags CarTypes
// @Accept json
// @Produce json
// @Param Request body dto.CreateCarTypeRequest true "Create a CarType"
// @Success 201 {object} dto.CarTypeResponse "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-types/ [post]
// @Security AuthBearer
func (h *CarTypeHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

//Update

// UpdateCarType godoc
// @Summary Update a CarType
// @Description Update a CarType
// @Tags CarTypes
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarTypeRequest true "Update a CarType"
// @Success 200 {object} dto.CarTypeResponse "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-types/{id} [put]
// @Security AuthBearer
func (h *CarTypeHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

//Deleted

// DeleteCarType godoc
// @Summary Delete a CarType
// @Description Delete a CarType
// @Tags CarTypes
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Deleted successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-types/{id} [delete]
// @Security AuthBearer
func (h *CarTypeHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

//get

// GetCarType godoc
// @Summary Get a CarType
// @Description Get a CarType
// @Tags CarTypes
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} dto.CarTypeResponse "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-types/{id} [get]
// @Security AuthBearer
func (h *CarTypeHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCarTypes godoc
// @Summary Filter Countries by pagination & criteria
// @Description Returns paginated list of countries based on filters
// @Tags CarTypes
// @Accept json
// @Produce json
// @Param Request body  dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-types/get-by-filter [post]
// @Security AuthBearer
func (h *CarTypeHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
