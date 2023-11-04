package service

import (
	"context"
	"time"
	"warrantyapi/model"
)

type PromotionService interface {
	Create(ctx context.Context, promotionInput []model.PromotionRequest, createdBy string) []model.PromotionResponse
	Update(ctx context.Context, promotionInput []model.PromotionRequest, updatedBy string) []model.PromotionResponse
	Delete(ctx context.Context, id string, deletedBy string) bool
	List(ctx context.Context, offset int, limit int, order string, searchRequest model.PromotionRequest) (responses []model.PromotionResponse)
	ListActivePromotion(ctx context.Context, offset int, limit int, order string, searchRequest model.PromotionRequest, warrantyCreated time.Time) []model.PromotionResponse
}
