package service

import (
	"context"
	"warrantyapi/model"
)

type WarrantyService interface {
	Create(ctx context.Context, warrantyRequest model.WarrantyRequest, createdBy string) model.WarrantyResponse
	FindById(ctx context.Context, id string) model.WarrantyResponse
	Update(ctx context.Context, warrantyRequest []model.WarrantyRequest, updatedBy string) []model.WarrantyResponse
	Delete(ctx context.Context, id string, deletedBy string) bool
	List(ctx context.Context, offset int, limit int, order string, warrantySearch model.WarrantyRequest) []model.WarrantyResponse
}
