package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
)

type PersianYearService struct {
	base *BaseService[models.PersianYear, dto.CreateYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]
}

func NewPersianYearService(cfg *config.Config) *PersianYearService {
	return &PersianYearService{
		base: &BaseService[models.PersianYear, dto.CreateYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),

		},
	}

}

func (s *PersianYearService) Create(ctx context.Context, req *dto.CreateYearRequest) (*dto.PersianYearResponse, error) {
	return s.base.Create(ctx, req)

}

func (s *PersianYearService) Update(ctx context.Context, id int, req *dto.UpdatePersianYearRequest) (*dto.PersianYearResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *PersianYearService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *PersianYearService) GetById(ctx context.Context, id int) (*dto.PersianYearResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *PersianYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PersianYearResponse], error) {

	return s.base.GetByFilter(ctx, req)
}
