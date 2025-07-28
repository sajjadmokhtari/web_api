package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"

	"github.com/gin-gonic/gin"
)

type PersianYearHandler struct {
	service *services.PersianYearService
}

func NewPersianYearHandler(cfg *config.Config) *PersianYearHandler {
	return &PersianYearHandler{
		service: services.NewPersianYearService(cfg),
	}
}

// CreatePersianYear godoc
// @Summary Create a PersianYear
// @Description Create a PersianYear
// @Tags PersianYears
// @Accept json
// @Produce json
// @Param Request body dto.CreateYearRequest true "Create a PersianYear"
// @Success 201 {object} dto.PersianYearResponse "PersianYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /years/ [post]
// @Security AuthBearer
func (h *PersianYearHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

//Update

// UpdatePersianYear godoc
// @Summary Update a PersianYear
// @Description Update a PersianYear
// @Tags PersianYears
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePersianYearRequest true "Update a PersianYear"
// @Success 200 {object} dto.PersianYearResponse "PersianYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /years/{id} [put]
// @Security AuthBearer
func (h *PersianYearHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

//Deleted

// DeletePersianYear godoc
// @Summary Delete a PersianYear
// @Description Delete a PersianYear
// @Tags PersianYears
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Deleted successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /years/{id} [delete]
// @Security AuthBearer
func (h *PersianYearHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

//get

// GetPersianYear godoc
// @Summary Get a PersianYear
// @Description Get a PersianYear
// @Tags PersianYears
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} dto.PersianYearResponse "PersianYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /years/{id} [get]
// @Security AuthBearer
func (h *PersianYearHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetPersianYears godoc
// @Summary Filter Countries by pagination & criteria
// @Description Returns paginated list of countries based on filters
// @Tags PersianYears
// @Accept json
// @Produce json
// @Param Request body  dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse "PersianYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /years/get-by-filter [post]
// @Security AuthBearer
func (h *PersianYearHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
