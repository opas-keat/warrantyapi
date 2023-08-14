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
func (repository *dealerRepositoryImpl) Insert(ctx context.Context, dealers entity.Dealer) entity.Dealer {
	err := repository.DB.WithContext(ctx).Create(&dealers).Error
	exception.PanicLogging(err)
	print(dealers.DealerName)
	return dealers
}
