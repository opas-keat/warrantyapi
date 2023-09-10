package repository

import (
	"context"
	"warrantyapi/entity"
)

type ProductRepository interface {
	Insert(ctx context.Context, products []entity.Product) []entity.Product
	GetById(ctx context.Context, id int) entity.Product
	Update(ctx context.Context, products []entity.Product) []entity.Product
	Delete(ctx context.Context, product entity.Product) bool
	List(ctx context.Context, offset int, limit int, order string, searchInput entity.Product) []entity.Product
	Total(ctx context.Context, searchInput entity.Product) int64
	CheckDuplicate(ctx context.Context) bool
}
