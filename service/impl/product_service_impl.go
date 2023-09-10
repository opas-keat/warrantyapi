package impl

import (
	"context"
	"warrantyapi/constant"
	"warrantyapi/entity"
	"warrantyapi/model"
	"warrantyapi/repository"
	"warrantyapi/service"
)

func NewProductServiceImpl(productRepository *repository.ProductRepository, logRepository *repository.LogRepository) service.ProductService {
	return &productServiceImpl{
		ProductRepository: *productRepository,
		LogRepository:     *logRepository,
	}
}

type productServiceImpl struct {
	repository.ProductRepository
	repository.LogRepository
}

// Create implements service.ProductService
func (service *productServiceImpl) Create(ctx context.Context, productInput []model.ProductRequest, createdBy string) []model.ProductResponse {
	var products []entity.Product
	for _, product := range productInput {
		products = append(products, entity.Product{
			CreatedBy:     createdBy,
			ProductType:   product.ProductType,
			ProductBrand:  product.ProductBrand,
			ProductAmount: product.ProductAmount,
			WarrantyNo:    product.WarrantyNo,
		})
	}
	service.ProductRepository.Insert(ctx, products)
	var responses []model.ProductResponse
	for _, rs := range products {
		service.LogRepository.Insert(ctx, entity.Log{
			CreatedBy: createdBy,
			Module:    constant.ModuleProduct,
			Detail:    "สร้าง : สินค้า รหัส  " + rs.ID.String() + " " + rs.ProductBrand,
		})

		responses = append(responses, model.ProductResponse{
			ID:            rs.ID.String(),
			ProductType:   rs.ProductType,
			ProductBrand:  rs.ProductBrand,
			ProductAmount: rs.ProductAmount,
			WarrantyNo:    rs.WarrantyNo,
		})
	}
	return responses
}

// List implements service.ProductService
func (service *productServiceImpl) List(ctx context.Context, offset int, limit int, order string, searchRequest model.ProductRequest) (responses []model.ProductResponse) {
	searchEntity := entity.Product{
		WarrantyNo: searchRequest.WarrantyNo,
	}
	results := service.ProductRepository.List(ctx, offset, limit, order, searchEntity)
	if len(results) == 0 {
		return []model.ProductResponse{}
	}

	for _, rs := range results {
		responses = append(responses, model.ProductResponse{
			ID:            rs.ID.String(),
			ProductType:   rs.ProductType,
			ProductBrand:  rs.ProductBrand,
			ProductAmount: rs.ProductAmount,
			WarrantyNo:    rs.WarrantyNo,
		})
	}
	return responses
}
