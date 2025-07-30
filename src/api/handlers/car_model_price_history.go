package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"

	"github.com/gin-gonic/gin"
)

type CarModelPriceHistoryHandler struct {
	service *services.CarModelPriceHistoryService
}

func NewCarModelPriceHistoryHandler(cfg *config.Config) *CarModelPriceHistoryHandler {
	return &CarModelPriceHistoryHandler{
		service: services.NewCarModelPriceHistoryService(cfg),
	}
}

// CreateCarModelPriceHistory godoc
// @Summary Create a CarModelPriceHistory
// @Description Create a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produce json
// @Param Request body dto.CreateCarModelPriceHistoryRequest true "Create a CarModelPriceHistory"
// @Success 201 {object} dto.CarModelPriceHistoryResponse "CarModelPriceHistory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-model-price-histories/ [post]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

//Update

// UpdateCarModelPriceHistory godoc
// @Summary Update a CarModelPriceHistory
// @Description Update a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelPriceHistoryRequest true "Update a CarModelPriceHistory"
// @Success 200 {object} dto.CarModelPriceHistoryResponse "CarModelPriceHistory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-model-price-histories/{id} [put]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

//Deleted

// DeleteCarModelPriceHistory godoc
// @Summary Delete a CarModelPriceHistory
// @Description Delete a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Deleted successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-model-price-histories/{id} [delete]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

//get

// GetCarModelPriceHistory godoc
// @Summary Get a CarModelPriceHistory
// @Description Get a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} dto.CarModelPriceHistoryResponse "CarModelPriceHistory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-model-price-histories/{id} [get]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCarModelPriceHistories godoc
// @Summary Filter Countries by pagination & criteria
// @Description Returns paginated list of countries based on filters
// @Tags CarModelPriceHistories
// @Accept json
// @Produce json
// @Param Request body  dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse "CarModelPriceHistory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-model-price-histories/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
