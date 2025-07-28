package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
)

type ColorService struct {
	base *BaseService[models.Color, dto.UpdateColorRequest, dto.UpdateColorRequest, dto.ColorResponse]
}

func NewColorService(cfg *config.Config) *ColorService {
	return &ColorService{
		base: &BaseService[models.Color, dto.UpdateColorRequest, dto.UpdateColorRequest, dto.ColorResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),

			
		},
	}

}

func (s *ColorService) Create(ctx context.Context, req *dto.UpdateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Create(ctx, req)

}

func (s *ColorService) Update(ctx context.Context, id int, req *dto.UpdateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *ColorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *ColorService) GetById(ctx context.Context, id int) (*dto.ColorResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *ColorService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.ColorResponse], error) {

	return s.base.GetByFilter(ctx, req)
}
