package impl

import (
	"context"
	"errors"
	"warrantyapi/constant"
	"warrantyapi/entity"
	"warrantyapi/repository"

	"gorm.io/gorm"
)

func NewAuthenRepositoryImpl(DB *gorm.DB) repository.AuthenRepository {
	return &authenRepositoryImpl{DB: DB}
}

type authenRepositoryImpl struct {
	*gorm.DB
}

func (repository *authenRepositoryImpl) Login(ctx context.Context, userName string) (entity.UserAuth, error) {
	var userAuth entity.UserAuth
	result := repository.DB.Model(&entity.User{}).WithContext(ctx).Unscoped().Where("user_name = ?", userName).Where("deleted_at is null").First(&userAuth)
	if result.RowsAffected == 0 {
		return entity.UserAuth{}, errors.New(constant.MESSAGE_LOGIN_FAIL)
	}
	return userAuth, nil
}

func (repository *authenRepositoryImpl) LogOut(ctx context.Context) {
	panic("unimplemented")
}
