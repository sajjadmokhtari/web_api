package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dto.CreateGearboxTypeRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]
}

func NewGearboxService(cfg *config.Config) *GearboxService {
	return &GearboxService{
		base: &BaseService[models.Gearbox, dto.CreateGearboxTypeRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),

		},
	}

}

func (s *GearboxService) Create(ctx context.Context, req *dto.CreateGearboxTypeRequest) (*dto.GearboxResponse, error) {
	return s.base.Create(ctx, req)

}

func (s *GearboxService) Update(ctx context.Context, id int, req *dto.UpdateGearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *GearboxService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *GearboxService) GetById(ctx context.Context, id int) (*dto.GearboxResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *GearboxService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GearboxResponse], error) {

	return s.base.GetByFilter(ctx, req)
}
