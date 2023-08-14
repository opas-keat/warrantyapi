package service

import (
	"context"
)

type LogService interface {
	Create(ctx context.Context, module string, detail string)
	// FindById(ctx context.Context, userId string) model.UserResponse
	// FindAll(ctx context.Context, limit int, offset int, orderBy string, sort string) []model.UserResponse
}
