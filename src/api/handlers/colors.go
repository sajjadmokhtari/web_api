package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"

	"github.com/gin-gonic/gin"
)

type ColorHandler struct {
	service *services.ColorService
}

func NewColorHandler(cfg *config.Config) *ColorHandler {
	return &ColorHandler{
		service: services.NewColorService(cfg),
	}
}

// CreateColor godoc
// @Summary Create a Color
// @Description Create a Color
// @Tags Colors
// @Accept json
// @Produce json
// @Param Request body dto.CreateColorRequest true "Create a Color"
// @Success 201 {object} dto.ColorResponse "Color response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /colors/ [post]
// @Security AuthBearer
func (h *ColorHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

//Update

// UpdateColor godoc
// @Summary Update a Color
// @Description Update a Color
// @Tags Colors
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param Request body dto.UpdateColorRequest true "Update a Color"
// @Success 200 {object} dto.ColorResponse "Color response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /colors/{id} [put]
// @Security AuthBearer
func (h *ColorHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

//Deleted

// DeleteColor godoc
// @Summary Delete a Color
// @Description Delete a Color
// @Tags Colors
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Deleted successfully"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /colors/{id} [delete]
// @Security AuthBearer
func (h *ColorHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

//get

// GetColor godoc
// @Summary Get a Color
// @Description Get a Color
// @Tags Colors
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} dto.ColorResponse "Color response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /colors/{id} [get]
// @Security AuthBearer
func (h *ColorHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// Getcolors godoc
// @Summary Filter Countries by pagination & criteria
// @Description Returns paginated list of countries based on filters
// @Tags Colors
// @Accept json
// @Produce json
// @Param Request body  dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse "Color response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /colors/get-by-filter [post]
// @Security AuthBearer
func (h *ColorHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
