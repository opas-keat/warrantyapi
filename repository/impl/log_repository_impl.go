package impl

import (
	"context"
	"warrantyapi/entity"
	"warrantyapi/exception"
	"warrantyapi/repository"

	"gorm.io/gorm"
)

func NewLogRepositoryImpl(DB *gorm.DB) repository.LogRepository {
	return &logRepositoryImpl{DB: DB}
}

type logRepositoryImpl struct {
	*gorm.DB
}

// Insert implements repository.LogRepository
func (repository *logRepositoryImpl) Insert(ctx context.Context, log entity.Log) entity.Log {
	err := repository.DB.WithContext(ctx).Create(&log).Error
	exception.PanicLogging(err)
	return log
}

// FindById implements repository.LogRepository
func (repository *logRepositoryImpl) FindById(ctx context.Context, logId string) (entity.Log, error) {
	panic("unimplemented")
}

// FindAll implements repository.LogRepository
func (repository *logRepositoryImpl) FindAll(ctx context.Context, limit int, offset int, orderBy string, sort string) []entity.Log {
	panic("unimplemented")
}
