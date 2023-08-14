package repository

import (
	"context"
	"warrantyapi/entity"
)

type DealerRepository interface {
	Insert(ctx context.Context, budgets entity.Dealer) entity.Dealer
}
