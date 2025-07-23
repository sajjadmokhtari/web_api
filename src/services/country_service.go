package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
)

type CountryService struct {
	base *BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
    return &CountryService{
        base: &BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]{
            Database: db.GetDb(),
            Logger:   logging.NewLogger(cfg),
            Preload:  []preload{{string: "Cities"}}, // ✨ این خط کلید حل مشکله
        },
    }
}


func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	return s.base.Create(ctx, req)

}

func (s *CountryService) Update(ctx context.Context, id int, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *CountryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *CountryService) GetById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *CountryService) GetByFilter(ctx context.Context,req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CountryResponse],error ) {

	return s.base.GetByFilter(ctx, req)
}
