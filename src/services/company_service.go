package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
)

type CompanyService struct {
	base *BaseService[models.Company, dto.UpdateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]
}

func NewCompanyService(cfg *config.Config) *CompanyService {
	return &CompanyService{
		base: &BaseService[models.Company, dto.UpdateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preload:  []preload{{string: "Country"}}, // 🌟 این خط کمپانی رو کامل می‌کنه
		},
	}
}

func (s *CompanyService) Create(ctx context.Context, req *dto.UpdateCompanyRequest) (*dto.CompanyResponse, error) {
	return s.base.Create(ctx, req)

}

func (s *CompanyService) Update(ctx context.Context, id int, req *dto.UpdateCompanyRequest) (*dto.CompanyResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *CompanyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *CompanyService) GetById(ctx context.Context, id int) (*dto.CompanyResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *CompanyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CompanyResponse], error) {

	return s.base.GetByFilter(ctx, req)
}
