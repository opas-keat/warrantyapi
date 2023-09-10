package service

import (
	"context"
	"warrantyapi/model"
)

type ProductService interface {
	Create(ctx context.Context, productInput []model.ProductRequest, createdBy string) []model.ProductResponse
	List(ctx context.Context, offset int, limit int, order string, searchRequest model.ProductRequest) (responses []model.ProductResponse)
	// FindById(ctx context.Context, stationId string) model.CommissResponse
}
