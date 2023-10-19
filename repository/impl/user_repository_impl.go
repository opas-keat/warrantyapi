package impl

import (
	"context"
	"errors"
	"warrantyapi/constant"
	"warrantyapi/entity"
	"warrantyapi/exception"
	"warrantyapi/repository"

	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

// Insert implements repository.UserRepository
func (repository *userRepositoryImpl) Insert(ctx context.Context, user entity.User) entity.User {
	err := repository.DB.WithContext(ctx).Create(&user).Error
	exception.PanicLogging(err)
	return user
}

// Update implements repository.UserRepository
func (repository *userRepositoryImpl) Update(ctx context.Context, user entity.User) entity.User {
	err := repository.DB.WithContext(ctx).Updates(&user).Error
	exception.PanicLogging(err)
	return user
}

// Delete implements repository.UserRepository
func (repository *userRepositoryImpl) Delete(ctx context.Context, user entity.User) {
	err := repository.DB.WithContext(ctx).Delete(&user).Error
	exception.PanicLogging(err)
}

// FindById implements repository.UserRepository
func (repository *userRepositoryImpl) FindById(ctx context.Context, userId string) (entity.User, error) {
	var user entity.User
	result := repository.DB.WithContext(ctx).Unscoped().Debug().Where("id = ?", userId).First(&user)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New(constant.MESSAGE_USER_NOT_FOUND)
	}
	return user, nil
}

// FindAll implements repository.UserRepository
func (repository *userRepositoryImpl) FindAll(ctx context.Context, limit int, offset int, orderBy string, sort string) []entity.User {
	var users []entity.User
	repository.DB.WithContext(ctx).Debug().
		// Table().
		// Where("").
		Offset(offset).
		Limit(limit).
		// Order("province desc").
		// Order("amphure desc").
		// Order("district desc").
		Find(&users)
	return users
}

// FindByUserName implements repository.UserRepository
func (repository *userRepositoryImpl) FindByUserName(ctx context.Context, userName string) (entity.User, error) {
	var user entity.User
	result := repository.DB.WithContext(ctx).Unscoped().Debug().Where("user_name = ?", userName).First(&user)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New(constant.MESSAGE_USER_NOT_FOUND)
	}
	return user, nil
}

// FindDuplicateUserName implements repository.UserRepository
func (repository *userRepositoryImpl) FindDuplicateUserName(ctx context.Context, userName string) (entity.User, error) {
	var user entity.User
	result := repository.DB.WithContext(ctx).Unscoped().Debug().Where("deleted_at is null and user_name = ?", userName).First(&user)
	if result.RowsAffected > 0 {
		return entity.User{}, errors.New(constant.MESSAGE_USER_DUPLICATE)
	}
	return user, nil
}
