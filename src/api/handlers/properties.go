package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"

	"github.com/gin-gonic/gin"
)

type PropertyHandler struct {
	service *services.PropertyService
}

func NewPropertyHandler(cfg *config.Config) *PropertyHandler {
	return &PropertyHandler{
		service: services.NewPropertyService(cfg),
	}
}

// CreatePropertyCategory godoc
// @Summary Create a PropertyCategory
// @Description Create a PropertyCategory
// @Tags Properties
// @Accept json
// @Produce json
// @Param Request body dto.CreatePropertyRequest true "Create a Property"
// @Success 201 {object} dto.PropertyCategoryResponse "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /properties/ [post]
// @Security AuthBearer
func (h *PropertyHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

//Update

// UpdatePropertyCategory godoc
// @Summary Update a PropertyCategory
// @Description Update a PropertyCategory
// @Tags Properties
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyRequest true "Update a Property"
// @Success 200 {object} dto.PropertyCategoryResponse "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /properties/{id} [put]
// @Security AuthBearer
func (h *PropertyHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

//Deleted

// DeletePropertyCategory godoc
// @Summary Delete a PropertyCategory
// @Description Delete a PropertyCategory
// @Tags Properties
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Deleted successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /properties/{id} [delete]
// @Security AuthBearer
func (h *PropertyHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

//get

// GetPropertyCategory godoc
// @Summary Get a PropertyCategory
// @Description Get a PropertyCategory
// @Tags Properties
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} dto.PropertyCategoryResponse "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /properties/{id} [get]
// @Security AuthBearer
func (h *PropertyHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetProperties godoc
// @Summary Filter Countries by pagination & criteria
// @Description Returns paginated list of countries based on filters
// @Tags Properties
// @Accept json
// @Produce json
// @Param Request body  dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /properties/get-by-filter [post]
// @Security AuthBearer
func (h *PropertyHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
