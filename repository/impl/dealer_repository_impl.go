package impl

import (
	"context"
	"warrantyapi/entity"
	"warrantyapi/exception"
	"warrantyapi/repository"

	"gorm.io/gorm"
)

func NewDealerRepositoryImpl(DB *gorm.DB) repository.DealerRepository {
	return &dealerRepositoryImpl{DB: DB}
}

type dealerRepositoryImpl struct {
	*gorm.DB
}

// Insert implements repository.DealerRepository
func (repository *dealerRepositoryImpl) Insert(ctx context.Context, dealers []entity.Dealer) []entity.Dealer {
	err := repository.DB.WithContext(ctx).Create(&dealers).Error
	exception.PanicLogging(err)
	return dealers
}

// List implements repository.DealerRepository
func (repository *dealerRepositoryImpl) List(ctx context.Context, offset int, limit int, order string, search entity.Dealer) []entity.Dealer {
	var result []entity.Dealer
	repository.DB.WithContext(ctx).Debug().
		Offset(offset).
		Limit(limit).
		Order(order).
		Where(search).
		Find(&result)
	return result
}
