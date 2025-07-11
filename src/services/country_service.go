package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/constants"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CountryService struct {
	database *gorm.DB
	logger   logging.Logger
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		database: db.GetDb(),
		logger:   logging.NewLogger(cfg),
	}
}
func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	fmt.Println("🟡 Create Country Called. Incoming Name:", req.Name)

	uidRaw := ctx.Value(constants.UserIdKey)
	fmt.Println("🔎 Raw User ID from context:", uidRaw)

	userID := int(uidRaw.(float64))
	fmt.Println("✅ Parsed User ID:", userID)

	country := models.Country{Name: req.Name}
	country.CreatedBy = userID
	country.CreatedAt = time.Now().UTC()

	fmt.Println("📦 Country model before DB insert:", country)

	tx := s.database.WithContext(ctx).Begin()
	fmt.Println("📡 Started transaction...")

	err := tx.Create(&country).Error
	if err != nil {
		fmt.Println("❌ Error during country insert:", err.Error())
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}

	tx.Commit()
	fmt.Println("✅ Transaction committed. Country ID:", country.Id)

	dto := &dto.CountryResponse{Name: country.Name, Id: country.Id}
	fmt.Println("🎯 Response DTO ready:", dto)

	return dto, nil
}







func (s *CountryService) Update(ctx context.Context, id int, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	updateMap := map[string]interface{}{
		"Name":        req.Name,
		"modified_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"modified_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}
	tx := s.database.WithContext(ctx).Begin()
	err := tx.
		Model(&models.Country{}).
		Where("id = ?", id).
		Updates(updateMap).
		Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}
	country := &models.Country{}
	err = tx.
		Model(&models.Country{}).
		Where("id = ? AND deleted_by is null", id).
		First(&country).
		Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err

	}
	tx.Commit()
	dto := &dto.CountryResponse{Name: country.Name, Id: country.Id}
	return dto, nil

}

func (s *CountryService) Delete(ctx context.Context, id int) error {
	tx := s.database.WithContext(ctx).Begin()
	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"deleted_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}
	if err := tx.
		Model(&models.Country{}).
		Where("id = ? ", id).
		Updates(deleteMap).
		Error; err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Delete, err.Error(), nil)
		return err

	}
	tx.Commit()

	return nil

}
func (s *CountryService) GetById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	country := &models.Country{}

	err := s.database.
		Where("id = ?", id).
		First(&country).
		Error
	if err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err

	}
	dto := &dto.CountryResponse{Name: country.Name, Id: country.Id}
	return dto, nil

}
