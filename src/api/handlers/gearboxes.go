package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"

	"github.com/gin-gonic/gin"
)

type GearboxHandler struct {
	service *services.GearboxService
}

func NewGearboxHandler(cfg *config.Config) *GearboxHandler {
	return &GearboxHandler{
		service: services.NewGearboxService(cfg),
	}
}

// CreateGearbox godoc
// @Summary Create a Gearbox
// @Description Create a Gearbox
// @Tags Gearboxes
// @Accept json
// @Produce json
// @Param Request body dto.CreateGearboxTypeRequest true "Create a Gearbox"
// @Success 201 {object} dto.GearboxResponse "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /gearboxes/ [post]
// @Security AuthBearer
func (h *GearboxHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

//Update

// UpdateGearbox godoc
// @Summary Update a Gearbox
// @Description Update a Gearbox
// @Tags Gearboxes
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param Request body dto.UpdateGearboxRequest true "Update a Gearbox"
// @Success 200 {object} dto.GearboxResponse "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /gearboxes/{id} [put]
// @Security AuthBearer
func (h *GearboxHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

//Deleted

// DeleteGearbox godoc
// @Summary Delete a Gearbox
// @Description Delete a Gearbox
// @Tags Gearboxes
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Deleted successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /gearboxes/{id} [delete]
// @Security AuthBearer
func (h *GearboxHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

//get

// GetGearbox godoc
// @Summary Get a Gearbox
// @Description Get a Gearbox
// @Tags Gearboxes
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} dto.GearboxResponse "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /gearboxes/{id} [get]
// @Security AuthBearer
func (h *GearboxHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetGearboxes godoc
// @Summary Filter Countries by pagination & criteria
// @Description Returns paginated list of countries based on filters
// @Tags Gearboxes
// @Accept json
// @Produce json
// @Param Request body  dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /gearboxes/get-by-filter [post]
// @Security AuthBearer
func (h *GearboxHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
