package service

import (
	"context"
	"warrantyapi/model"
)

type ConfigService interface {
	// Create(ctx context.Context, configInput []model.ConfigRequest, createdBy string) []model.ConfigResponse
	List(ctx context.Context, offset int, limit int, order string, searchRequest model.ConfigRequest) (responses []model.ConfigResponse)
	// FindById(ctx context.Context, stationId string) model.CommissResponse
}
