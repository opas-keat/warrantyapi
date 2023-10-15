package repository

import (
	"context"
	"warrantyapi/entity"
)

type ConfigRepository interface {
	// Insert(ctx context.Context, configs []entity.Config) []entity.Config
	GetById(ctx context.Context, id int) entity.Config
	Update(ctx context.Context, configs []entity.Config) []entity.Config
	// Delete(ctx context.Context, config entity.Config) bool
	List(ctx context.Context, offset int, limit int, order string, searchInput entity.Config) []entity.Config
	// Total(ctx context.Context, searchInput entity.Config) int64
	// CheckDuplicate(ctx context.Context) bool
}
