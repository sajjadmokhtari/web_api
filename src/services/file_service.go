package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
)

type FileService struct {
	base *BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse]
}

func NewFileService(cfg *config.Config) *FileService {
	return &FileService{
		base: &BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

func (s *FileService) Create(ctx context.Context, req *dto.CreateFileRequest) (*dto.FileResponse, error) {
	return s.base.Create(ctx, req)

}

func (s *FileService) Update(ctx context.Context, id int, req *dto.UpdateFileRequest) (*dto.FileResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *FileService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *FileService) GetById(ctx context.Context, id int) (*dto.FileResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *FileService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.FileResponse], error) {

	return s.base.GetByFilter(ctx, req)
}
