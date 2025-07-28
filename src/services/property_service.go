package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
)

type PropertyService struct {
	base *BaseService[models.Property, dto.UpdatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]
}

func NewPropertyService(cfg *config.Config) *PropertyService {
	return &PropertyService{
		base: &BaseService[models.Property, dto.UpdatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preload:  []preload{{string: "Category"}},
		},
	}
}

func (s *PropertyService) Create(ctx context.Context, req *dto.UpdatePropertyRequest) (*dto.PropertyResponse, error) {
	return s.base.Create(ctx, req)

}

func (s *PropertyService) Update(ctx context.Context, id int, req *dto.UpdatePropertyRequest) (*dto.PropertyResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *PropertyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *PropertyService) GetById(ctx context.Context, id int) (*dto.PropertyResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *PropertyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PropertyResponse], error) {

	return s.base.GetByFilter(ctx, req)
}
