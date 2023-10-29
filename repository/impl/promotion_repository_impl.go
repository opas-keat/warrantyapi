package impl

import (
	"context"
	"time"
	"warrantyapi/entity"
	"warrantyapi/exception"
	"warrantyapi/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewPromotionRepositoryImpl(DB *gorm.DB) repository.PromotionRepository {
	return &promotionRepositoryImpl{DB: DB}
}

type promotionRepositoryImpl struct {
	*gorm.DB
}

// Insert implements repository.PromotionRepository
func (repository *promotionRepositoryImpl) Insert(ctx context.Context, promotions []entity.Promotion) []entity.Promotion {
	err := repository.DB.WithContext(ctx).Create(&promotions).Error
	exception.PanicLogging(err)
	return promotions
}

// GetById implements repository.PromotionRepository
func (repository *promotionRepositoryImpl) GetById(ctx context.Context, id string) entity.Promotion {
	result := entity.Promotion{
		ID: uuid.MustParse(id),
	}
	repository.DB.WithContext(ctx).Debug().
		First(&result)
	return result
}

// Update implements repository.PromotionRepository
func (repository *promotionRepositoryImpl) Update(ctx context.Context, promotion []entity.Promotion) []entity.Promotion {
	err := repository.DB.WithContext(ctx).Save(&promotion).Error
	exception.PanicLogging(err)
	return promotion
}

// Delete implements repository.PromotionRepository
func (repository *promotionRepositoryImpl) Delete(ctx context.Context, promotion entity.Promotion) bool {
	repository.DB.WithContext(ctx).Debug().Delete(&promotion)
	return true
}

// List implements repository.PromotionRepository
func (repository *promotionRepositoryImpl) List(ctx context.Context, offset int, limit int, order string, search entity.Promotion) []entity.Promotion {
	var result []entity.Promotion
	repository.DB.WithContext(ctx).Debug().
		Offset(offset).
		Limit(limit).
		Order(order).
		Where(search).
		Find(&result)
	return result
}

// Total implements repository.PromotionRepository
func (repository *promotionRepositoryImpl) Total(ctx context.Context, search entity.Promotion) int64 {
	var count int64
	repository.DB.WithContext(ctx).Debug().
		Model(&entity.Promotion{}).
		Where(search).
		Count(&count)
	return count
}

// CheckDuplicate implements repository.PromotionRepository
func (repository *promotionRepositoryImpl) CheckDuplicate(ctx context.Context) bool {
	panic("unimplemented")
}

func (repository *promotionRepositoryImpl) ListActivePromotion(
	ctx context.Context,
	offset int, limit int, order string,
	search entity.Promotion,
	warrantyCreated time.Time,
) []entity.Promotion {
	var result []entity.Promotion
	repository.DB.WithContext(ctx).Debug().
		Offset(offset).
		Limit(limit).
		Order(order).
		Where(search).
		Where("promotion_from <= ?", warrantyCreated).
		Where("promotion_to >= ?", warrantyCreated).
		Find(&result)
	return result
}
