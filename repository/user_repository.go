package repository

import (
	"context"
	"warrantyapi/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user entity.User) entity.User
	Update(ctx context.Context, user entity.User) entity.User
	Delete(ctx context.Context, user entity.User)
	FindById(ctx context.Context, userId string) (entity.User, error)
	FindAll(ctx context.Context, limit int, offset int, orderBy string, sort string) []entity.User
	FindByUserName(ctx context.Context, userName string) (entity.User, error)
	FindDuplicateUserName(ctx context.Context, userName string) (entity.User, error)
}
