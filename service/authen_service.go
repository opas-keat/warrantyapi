package service

import (
	"context"
	"warrantyapi/entity"
	"warrantyapi/model"
)

type AuthenService interface {
	Login(ctx context.Context, userName string, passWord string) entity.UserAuth
	LogOut(ctx context.Context) model.AuthResponse
}
