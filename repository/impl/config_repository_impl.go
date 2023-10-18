package impl

import (
	"context"
	"warrantyapi/entity"
	"warrantyapi/exception"
	"warrantyapi/repository"

	"gorm.io/gorm"
)

func NewConfigRepositoryImpl(DB *gorm.DB) repository.ConfigRepository {
	return &configRepositoryImpl{DB: DB}
}

type configRepositoryImpl struct {
	*gorm.DB
}

// Insert implements repository.ConfigRepository
func (repository *configRepositoryImpl) Insert(ctx context.Context, configs []entity.Config) []entity.Config {
	err := repository.DB.WithContext(ctx).Create(&configs).Error
	exception.PanicLogging(err)
	return configs
}

// GetById implements repository.ConfigRepository
func (repository *configRepositoryImpl) GetById(ctx context.Context, id int) entity.Config {
	var result entity.Config
	repository.DB.WithContext(ctx).Debug().
		First(&result, id)
	return result
}

// Update implements repository.ConfigRepository
func (repository *configRepositoryImpl) Update(ctx context.Context, config []entity.Config) []entity.Config {
	err := repository.DB.WithContext(ctx).Save(&config).Error
	exception.PanicLogging(err)
	return config
}

// // UpdateWithConfigCode implements repository.ConfigRepository
// func (repository *configRepositoryImpl) UpdateWithConfigCode(ctx context.Context, config []entity.Config) []entity.Config {
// 	repository.DB.WithContext(ctx).Where("config_code = ?", ).Update("name", "hello")
// 	err := repository.DB.WithContext(ctx).Save(&config).Error
// 	exception.PanicLogging(err)
// 	return config
// }

// Delete implements repository.ConfigRepository
func (repository *configRepositoryImpl) Delete(ctx context.Context, config entity.Config) bool {
	repository.DB.WithContext(ctx).Debug().Delete(&config)
	return true
}

// List implements repository.ConfigRepository
func (repository *configRepositoryImpl) List(ctx context.Context, offset int, limit int, order string, search entity.Config) []entity.Config {
	var result []entity.Config
	repository.DB.WithContext(ctx).Debug().
		Offset(offset).
		Limit(limit).
		Order(order).
		Where(search).
		Find(&result)
	return result
}

// Total implements repository.ConfigRepository
func (repository *configRepositoryImpl) Total(ctx context.Context, search entity.Config) int64 {
	var count int64
	repository.DB.WithContext(ctx).Debug().
		Model(&entity.Config{}).
		Where(search).
		Count(&count)
	return count
}

// CheckDuplicate implements repository.ConfigRepository
func (repository *configRepositoryImpl) CheckDuplicate(ctx context.Context) bool {
	panic("unimplemented")
}
