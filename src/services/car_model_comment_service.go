package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/constants"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"context"
	"fmt"
	"log"
)

type CarModelCommentService struct {
	base *BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse]
}

func NewCarModelCommentService(cfg *config.Config) *CarModelCommentService {
	return &CarModelCommentService{
		base: &BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preload: []preload{
				{string: "User"},
			},
		},
	}

}

func (s *CarModelCommentService) Create(ctx context.Context, req *dto.CreateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	log.Println("🧪 CreateCarModelComment - Started")
	log.Printf("➡️ Incoming Request: CarModelId=%d, Message=%s", req.CarModelId, req.Message)

	
	rawUserId := ctx.Value(constants.UserIdKey)
	userFloat, ok := rawUserId.(float64)
	if !ok {
		log.Printf("❌ Invalid UserId in context: %v (type: %T)", rawUserId, rawUserId)
		return nil, fmt.Errorf("invalid user ID in context")
	}
	req.UserId = int(userFloat)

	log.Printf("✅ Final UserId set in request: %d", req.UserId)

	// اجرای عملیات ایجاد کامنت
	comment, err := s.base.Create(ctx, req)
	if err != nil {
		log.Println("❌ Error creating CarModelComment:", err)
		return nil, err
	}

	log.Printf("🎯 Comment created successfully with ID: %d", comment.Id)
	return comment, nil
}

func (s *CarModelCommentService) Update(ctx context.Context, id int, req *dto.UpdateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	return s.base.Update(ctx, req, id)

}

func (s *CarModelCommentService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)

}

func (s *CarModelCommentService) GetById(ctx context.Context, id int) (*dto.CarModelCommentResponse, error) {
	return s.base.GetById(ctx, id)
}
func (s *CarModelCommentService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelCommentResponse], error) {

	return s.base.GetByFilter(ctx, req)
}
