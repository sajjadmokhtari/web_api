package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
)

type CarModelImageService struct {
	base *BaseService[models.CarModelImage, dto.CreateCarModelImageRequest, dto.UpdateCarModelImageRequest, dto.CarModelImageResponse]
}

func NewCarModelImageService(cfg *config.Config) *CarModelImageService {
	return &CarModelImageService{
		base: &BaseService[models.CarModelImage, dto.CreateCarModelImageRequest, dto.UpdateCarModelImageRequest, dto.CarModelImageResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preload: []preload{
				{string: "Image"},
			},
		},
	}

}

func (s *CarModelImageService) Create(ctx context.Context, req *dto.CreateCarModelImageRequest) (*dto.CarModelImageResponse, error) {
	return s.base.Create(ctx, req)

}

func (s *CarModelImageService) Update(ctx context.Context, id int, req *dto.UpdateCarModelImageRequest) (*dto.CarModelImageResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *CarModelImageService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *CarModelImageService) GetById(ctx context.Context, id int) (*dto.CarModelImageResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *CarModelImageService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelImageResponse], error) {

	return s.base.GetByFilter(ctx, req)
}
