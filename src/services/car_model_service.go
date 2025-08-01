package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
)

type CarModelService struct {
	base *BaseService[models.CarModel, dto.UpdateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse]
}

func NewCarModelService(cfg *config.Config) *CarModelService {
	return &CarModelService{
		base: &BaseService[models.CarModel, dto.UpdateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preload: []preload{
				{string: "Company.Country"},
				{string: "CarType"},
				{string: "Gearbox"},
				{string: "CarModelColors.Color"},
				{string: "CarModelYears.PersianYear"},
				{string: "CarModelYears.CarModelPriceHistories"},
				{string: "CarModelImages.Image"},
			},
		},
	}

}

func (s *CarModelService) Create(ctx context.Context, req *dto.UpdateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Create(ctx, req)

}

func (s *CarModelService) Update(ctx context.Context, id int, req *dto.UpdateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *CarModelService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *CarModelService) GetById(ctx context.Context, id int) (*dto.CarModelResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *CarModelService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelResponse], error) {

	return s.base.GetByFilter(ctx, req)
}
