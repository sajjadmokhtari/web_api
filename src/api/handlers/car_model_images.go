package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"

	"github.com/gin-gonic/gin"
)

type CarModelImageHandler struct {
	service *services.CarModelImageService
}

func NewCarModelImageHandler(cfg *config.Config) *CarModelImageHandler {
	return &CarModelImageHandler{
		service: services.NewCarModelImageService(cfg),
	}
}

// CreateCarModelImage godoc
// @Summary Create a CarModelImage
// @Description Create a CarModelImage
// @Tags CarModelImages
// @Accept json
// @Produce json
// @Param Request body dto.CreateCarModelImageRequest true "Create a CarModelImage"
// @Success 201 {object} dto.CarModelImageResponse "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-model-images/ [post]
// @Security AuthBearer
func (h *CarModelImageHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

//Update

// UpdateCarModelImage godoc
// @Summary Update a CarModelImage
// @Description Update a CarModelImage
// @Tags CarModelImages
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelImageRequest true "Update a CarModelImage"
// @Success 200 {object} dto.CarModelImageResponse "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-model-images/{id} [put]
// @Security AuthBearer
func (h *CarModelImageHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

//Deleted

// DeleteCarModelImage godoc
// @Summary Delete a CarModelImage
// @Description Delete a CarModelImage
// @Tags CarModelImages
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Deleted successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-model-images/{id} [delete]
// @Security AuthBearer
func (h *CarModelImageHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

//get

// GetCarModelImage godoc
// @Summary Get a CarModelImage
// @Description Get a CarModelImage
// @Tags CarModelImages
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} dto.CarModelImageResponse "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-model-images/{id} [get]
// @Security AuthBearer
func (h *CarModelImageHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCarModelImages godoc
// @Summary Filter Countries by pagination & criteria
// @Description Returns paginated list of countries based on filters
// @Tags CarModelImages
// @Accept json
// @Produce json
// @Param Request body  dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /car-model-images/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelImageHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
