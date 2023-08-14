package service

import (
	"context"
	"warrantyapi/model"
)

type DealerService interface {
	Create(ctx context.Context, dealerInput model.DealerRequest, createdBy string) model.DealerResponse
	// FindById(ctx context.Context, stationId string) model.CommissResponse
	// List(ctx context.Context, offset int, limit int, order string, dealer model.Dealer) []model.DealerResponse
}
