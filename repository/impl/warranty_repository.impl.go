package impl

import (
	"context"
	"warrantyapi/entity"
	"warrantyapi/exception"
	"warrantyapi/repository"

	"gorm.io/gorm"
)

func NewWarrantyRepositoryImpl(DB *gorm.DB) repository.WarrantyRepository {
	return &warrantyRepositoryImpl{DB: DB}
}

type warrantyRepositoryImpl struct {
	*gorm.DB
}

// Insert implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) Insert(ctx context.Context, warrantys []entity.Warranty) []entity.Warranty {
	err := repository.DB.WithContext(ctx).Create(&warrantys).Error
	exception.PanicLogging(err)
	return warrantys
}

// GetById implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) GetById(ctx context.Context, id int) entity.Warranty {
	var result entity.Warranty
	repository.DB.WithContext(ctx).Debug().
		First(&result, id)
	return result
}

// Update implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) Update(ctx context.Context, warranty []entity.Warranty) []entity.Warranty {
	err := repository.DB.WithContext(ctx).Save(&warranty).Error
	exception.PanicLogging(err)
	return warranty
}

// Delete implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) Delete(ctx context.Context, warranty entity.Warranty) bool {
	repository.DB.WithContext(ctx).Debug().Delete(&warranty)
	return true
}

// List implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) List(ctx context.Context, offset int, limit int, order string, search entity.Warranty) []entity.Warranty {
	var result []entity.Warranty
	repository.DB.WithContext(ctx).Debug().
		Offset(offset).
		Limit(limit).
		Order(order).
		Where(search).
		Find(&result)
	return result
}

// Total implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) Total(ctx context.Context, search entity.Warranty) int64 {
	var count int64
	repository.DB.WithContext(ctx).Debug().
		Model(&entity.Warranty{}).
		Where(search).
		Count(&count)
	return count
}

// CheckDuplicate implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) CheckDuplicate(ctx context.Context) bool {
	panic("unimplemented")
}
