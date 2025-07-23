package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
)

type CityService struct {
	base *BaseService[models.City, dto.CreateUpdateCityRequest, dto.CreateUpdateCityRequest, dto.CityResponse]
}

func NewCityService(cfg *config.Config) *CityService {
	return &CityService{
		base: &BaseService[models.City, dto.CreateUpdateCityRequest, dto.CreateUpdateCityRequest, dto.CityResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preload: []preload{
				{string: "Country"},
			},
		},
	}

}

func (s *CityService) Create(ctx context.Context, req *dto.CreateUpdateCityRequest) (*dto.CityResponse, error) {
	return s.base.Create(ctx, req)

}

func (s *CityService) Update(ctx context.Context, id int, req *dto.CreateUpdateCityRequest) (*dto.CityResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *CityService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *CityService) GetById(ctx context.Context, id int) (*dto.CityResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *CityService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CityResponse], error) {

	return s.base.GetByFilter(ctx, req)
}
