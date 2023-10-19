package service

import (
	"context"
	"warrantyapi/model"
)

type UserService interface {
	Create(ctx context.Context, userInput model.User, createdBy string) model.UserResponse
	Update(ctx context.Context, userInput model.User, updatedBy string) model.UserResponse
	Delete(ctx context.Context, userId string, createdBy string)
	FindById(ctx context.Context, userId string) model.UserResponse
	FindAll(ctx context.Context, limit int, offset int, orderBy string, sort string) []model.UserResponse
	FindByUserName(ctx context.Context, userName string) model.UserResponse
}
