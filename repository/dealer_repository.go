package repository

import (
	"context"
	"warrantyapi/entity"
)

type DealerRepository interface {
	Insert(ctx context.Context, dealers []entity.Dealer) []entity.Dealer
	List(ctx context.Context, offset int, limit int, order string, search entity.Dealer) []entity.Dealer
}
