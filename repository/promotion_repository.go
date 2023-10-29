package repository

import (
	"context"
	"time"
	"warrantyapi/entity"
)

type PromotionRepository interface {
	Insert(ctx context.Context, promotions []entity.Promotion) []entity.Promotion
	GetById(ctx context.Context, id string) entity.Promotion
	Update(ctx context.Context, promotions []entity.Promotion) []entity.Promotion
	Delete(ctx context.Context, promotion entity.Promotion) bool
	List(ctx context.Context, offset int, limit int, order string, searchInput entity.Promotion) []entity.Promotion
	Total(ctx context.Context, searchInput entity.Promotion) int64
	CheckDuplicate(ctx context.Context) bool
	ListActivePromotion(ctx context.Context, offset int, limit int, order string, search entity.Promotion, warrantyCreated time.Time) []entity.Promotion
}
