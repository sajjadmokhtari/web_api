package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
)

type CarModelYearService struct {
	base *BaseService[models.CarModelYear, dto.CreateCarModelYearRequest, dto.UpdateCarModelYearRequest, dto.CarModelYearResponse]
}

func NewCarModelYearService(cfg *config.Config) *CarModelYearService {
	return &CarModelYearService{
		base: &BaseService[models.CarModelYear, dto.CreateCarModelYearRequest, dto.UpdateCarModelYearRequest, dto.CarModelYearResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preload: []preload{
				{string: "PersianYear"},
			},
		},
	}
}

func (s *CarModelYearService) Create(ctx context.Context, req *dto.CreateCarModelYearRequest) (*dto.CarModelYearResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelYearService) Update(ctx context.Context, id int, req *dto.UpdateCarModelYearRequest) (*dto.CarModelYearResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *CarModelYearService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *CarModelYearService) GetById(ctx context.Context, id int) (*dto.CarModelYearResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *CarModelYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelYearResponse], error) {

	return s.base.GetByFilter(ctx, req)
}
