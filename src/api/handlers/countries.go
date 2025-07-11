package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/api/helper"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"
	"net/http"
	"strconv"

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
	req := dto.CreateUpdateCountryRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	res, err := h.service.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return

	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
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
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := dto.CreateUpdateCountryRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	res, err := h.service.Update(c, id, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return

	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
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
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	err := h.service.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, 121, err))

		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, 0))
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
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	res, err := h.service.GetById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return

	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

//get by id




func (h *CountryHandler) GetByFilter(c *gin.Context) {}

