package impl

import (
	"context"
	"warrantyapi/entity"
	"warrantyapi/exception"
	"warrantyapi/repository"

	"gorm.io/gorm"
)

func NewProductRepositoryImpl(DB *gorm.DB) repository.ProductRepository {
	return &productRepositoryImpl{DB: DB}
}

type productRepositoryImpl struct {
	*gorm.DB
}

// Insert implements repository.ProductRepository
func (repository *productRepositoryImpl) Insert(ctx context.Context, products []entity.Product) []entity.Product {
	err := repository.DB.WithContext(ctx).Create(&products).Error
	exception.PanicLogging(err)
	return products
}

// GetById implements repository.ProductRepository
func (repository *productRepositoryImpl) GetById(ctx context.Context, id int) entity.Product {
	var result entity.Product
	repository.DB.WithContext(ctx).Debug().
		First(&result, id)
	return result
}

// Update implements repository.ProductRepository
func (repository *productRepositoryImpl) Update(ctx context.Context, product []entity.Product) []entity.Product {
	err := repository.DB.WithContext(ctx).Save(&product).Error
	exception.PanicLogging(err)
	return product
}

// Delete implements repository.ProductRepository
func (repository *productRepositoryImpl) Delete(ctx context.Context, product entity.Product) bool {
	repository.DB.WithContext(ctx).Debug().Delete(&product)
	return true
}

// List implements repository.ProductRepository
func (repository *productRepositoryImpl) List(ctx context.Context, offset int, limit int, order string, search entity.Product) []entity.Product {
	var result []entity.Product
	repository.DB.WithContext(ctx).Debug().
		Offset(offset).
		Limit(limit).
		Order(order).
		Where(search).
		Find(&result)
	return result
}

// Total implements repository.ProductRepository
func (repository *productRepositoryImpl) Total(ctx context.Context, search entity.Product) int64 {
	var count int64
	repository.DB.WithContext(ctx).Debug().
		Model(&entity.Product{}).
		Where(search).
		Count(&count)
	return count
}

// CheckDuplicate implements repository.ProductRepository
func (repository *productRepositoryImpl) CheckDuplicate(ctx context.Context) bool {
	panic("unimplemented")
}
