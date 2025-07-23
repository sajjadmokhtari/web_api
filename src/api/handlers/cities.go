package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"

	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{
		service: services.NewCityService(cfg),
	}
}

// CreateCity godoc
// @Summary Create a City
// @Description Create a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param Request body dto.CreateUpdateCityRequest true "Create a City"
// @Success 201 {object} dto.CityResponse "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /Cities/ [post]
// @Security AuthBearer
func (h *CityHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

//Update

// UpdateCity godoc
// @Summary Update a City
// @Description Update a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdateCityRequest true "Update a City"
// @Success 200 {object} dto.CityResponse "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /Cities/{id} [put]
// @Security AuthBearer
func (h *CityHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

//Deleted

// DeleteCity godoc
// @Summary Delete a City
// @Description Delete a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Deleted successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /Cities/{id} [delete]
// @Security AuthBearer
func (h *CityHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

//get

// GetCity godoc
// @Summary Get a City
// @Description Get a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} dto.CityResponse "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /Cities/{id} [get]
// @Security AuthBearer
func (h *CityHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCities godoc
// @Summary Filter Countries by pagination & criteria
// @Description Returns paginated list of countries based on filters
// @Tags Cities
// @Accept json
// @Produce json
// @Param Request body  dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /Cities/get-by-filter [post]
// @Security AuthBearer
func (h *CityHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
