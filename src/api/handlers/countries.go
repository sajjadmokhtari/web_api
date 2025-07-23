package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"

	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{service: services.NewCountryService(cfg)}
}

// CreateCountry godoc
// @Summary Create a country
// @Description Create a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param Request body dto.CreateUpdateCountryRequest true "Create a country"
// @Success 201 {object} dto.CountryResponse "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /countries/ [post]
// @Security AuthBearer
func (h *CountryHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

//Update

// UpdateCountry godoc
// @Summary Update a country
// @Description Update a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdateCountryRequest true "Update a country"
// @Success 200 {object} dto.CountryResponse "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /countries/{id} [put]
// @Security AuthBearer
func (h *CountryHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

//Deleted

// DeleteCountry godoc
// @Summary Delete a country
// @Description Delete a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Deleted successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /countries/{id} [delete]
// @Security AuthBearer
func (h *CountryHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)

}

//get

// GetCountry godoc
// @Summary Get a country
// @Description Get a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} dto.CountryResponse "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /countries/{id} [get]
// @Security AuthBearer
func (h *CountryHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCountries godoc
// @Summary Filter Countries by pagination & criteria
// @Description Returns paginated list of countries based on filters
// @Tags Countries
// @Accept json
// @Produce json
// @Param Request body  dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /countries/get-by-filter [post]
// @Security AuthBearer
func (h *CountryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
