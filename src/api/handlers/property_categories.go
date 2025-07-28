package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"

	"github.com/gin-gonic/gin"
)

type PropertyCategoryHandler struct {
	service *services.PropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	return &PropertyCategoryHandler{
		service: services.NewPropertyCategoryService(cfg),
	}
}

// CreatePropertyCategory godoc
// @Summary Create a PropertyCategory
// @Description Create a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @Produce json
// @Param Request body dto.CreatePropertyCategoryRequest true "Create a PropertyCategory"
// @Success 201 {object} dto.PropertyCategoryResponse "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /property-categories/ [post]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

//Update

// UpdatePropertyCategory godoc
// @Summary Update a PropertyCategory
// @Description Update a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyCategoryRequest true "Update a PropertyCategory"
// @Success 200 {object} dto.PropertyCategoryResponse "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /property-categories/{id} [put]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

//Deleted

// DeletePropertyCategory godoc
// @Summary Delete a PropertyCategory
// @Description Delete a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Deleted successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /property-categories/{id} [delete]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

//get

// GetPropertyCategory godoc
// @Summary Get a PropertyCategory
// @Description Get a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} dto.PropertyCategoryResponse "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /property-categories/{id} [get]
// @Security AuthBearer
func (h *PropertyCategoryHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetPropertyCategories godoc
// @Summary Filter Countries by pagination & criteria
// @Description Returns paginated list of countries based on filters
// @Tags PropertyCategories
// @Accept json
// @Produce json
// @Param Request body  dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /property-categories/get-by-filter [post]
// @Security AuthBearer
func (h *PropertyCategoryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
