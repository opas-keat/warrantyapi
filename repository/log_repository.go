package repository

import (
	"context"
	"warrantyapi/entity"
)

type LogRepository interface {
	Insert(ctx context.Context, log entity.Log) entity.Log
	FindById(ctx context.Context, logId string) (entity.Log, error)
	FindAll(ctx context.Context, limit int, offset int, orderBy string, sort string) []entity.Log
}
