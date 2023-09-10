package service

import (
	"context"
	"warrantyapi/model"
)

type DealerService interface {
	Create(ctx context.Context, dealerInput []model.DealerRequest, createdBy string) []model.DealerResponse
	List(ctx context.Context, offset int, limit int, order string, searchRequest model.DealerRequest) (responses []model.DealerResponse)
	FindById(ctx context.Context, id string) []model.DealerResponse
}
