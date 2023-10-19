package repository

import (
	"context"
	"warrantyapi/entity"
)

type AuthenRepository interface {
	Login(ctx context.Context, userName string) (entity.UserAuth, error)
	LogOut(ctx context.Context)
}
