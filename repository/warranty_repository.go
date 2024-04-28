package repository

import (
	"context"
	"warrantyapi/entity"
)

type WarrantyRepository interface {
	Insert(ctx context.Context, warrantys []entity.Warranty) []entity.Warranty
	GetById(ctx context.Context, id string) entity.Warranty
	Update(ctx context.Context, warrantys []entity.Warranty) []entity.Warranty
	Delete(ctx context.Context, warranty entity.Warranty) bool
	List(ctx context.Context, offset int, limit int, order string, searchInput entity.Warranty) []entity.Warranty
	Total(ctx context.Context, searchInput entity.Warranty) int64
	CheckDuplicate(ctx context.Context) bool
	ListCustomer(ctx context.Context, offset int, limit int, order string, searchInput entity.Warranty) []entity.Warranty
	ListExcels(ctx context.Context, offset int, limit int, order string, searchInput entity.Warranty) []entity.Excels
}
