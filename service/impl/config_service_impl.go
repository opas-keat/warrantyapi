package impl

import (
	"context"
	"warrantyapi/entity"
	"warrantyapi/model"
	"warrantyapi/repository"
	"warrantyapi/service"
)

func NewConfigServiceImpl(
	configRepository *repository.ConfigRepository,
	logRepository *repository.LogRepository,
) service.ConfigService {
	return &configServiceImpl{
		ConfigRepository: *configRepository,
		LogRepository:    *logRepository,
	}
}

type configServiceImpl struct {
	repository.ConfigRepository
	repository.LogRepository
}

// // FindById implements service.ConfigService
// func (service *configServiceImpl) FindById(ctx context.Context, id string) model.WarrantyResponse {
// 	rs := service.WarrantyRepository.GetById(ctx, id)
// 	// var responses model.WarrantyResponse
// 	responses := model.WarrantyResponse{
// 		ID:                   rs.ID.String(),
// 		WarrantyNo:           rs.WarrantyNo,
// 		WarrantyDateTime:     rs.WarrantyDateTime,
// 		DealerCode:           rs.DealerCode,
// 		DealerName:           rs.DealerName,
// 		CustomerName:         rs.CustomerName,
// 		CustomerPhone:        rs.CustomerPhone,
// 		CustomerLicensePlate: rs.CustomerLicensePlate,
// 		CustomerEmail:        rs.CustomerEmail,
// 		CustomerMile:         rs.CustomerMile,
// 	}
// 	ProductSearch := entity.Product{
// 		WarrantyNo: rs.WarrantyNo,
// 	}
// 	responseProducts := service.ProductRepository.List(ctx, 0, 100, "product_type desc", ProductSearch)
// 	for _, responseProduct := range responseProducts {
// 		responses.ProductResponse = append(responses.ProductResponse, model.ProductResponse{
// 			ID:                     responseProduct.ID.String(),
// 			ProductType:            responseProduct.ProductType,
// 			ProductBrand:           responseProduct.ProductBrand,
// 			ProductAmount:          responseProduct.ProductAmount,
// 			ProductStructureExpire: responseProduct.ProductStructureExpire,
// 			ProductColorExpire:     responseProduct.ProductColorExpire,
// 			ProductTireExpire:      responseProduct.ProductTireExpire,
// 			ProductMileExpire:      responseProduct.ProductMileExpire,
// 			ProductPromotionExpire: responseProduct.ProductPromotionExpire,
// 			WarrantyNo:             responseProduct.WarrantyNo,
// 		})
// 	}
// 	return responses
// }

// // Update implements service.WarrantyService
// func (service *configServiceImpl) Update(ctx context.Context, warrantyInput []model.WarrantyRequest, updatedBy string) []model.WarrantyResponse {
// 	var warrantys []entity.Warranty
// 	for _, warranty := range warrantyInput {
// 		warrantys = append(warrantys, entity.Warranty{
// 			ID: uuid.MustParse(warranty.ID),
// 		})
// 	}
// 	service.WarrantyRepository.Update(ctx, warrantys)
// 	var responses []model.WarrantyResponse
// 	for _, rs := range warrantys {
// 		service.LogRepository.Insert(ctx, entity.Log{
// 			CreatedBy: updatedBy,
// 			Module:    constant.ModuleWarranty,
// 			Detail:    "แก้ไข : การรับประกัน รหัส  " + rs.ID.String() + " " + rs.WarrantyNo,
// 		})

// 		responses = append(responses, model.WarrantyResponse{
// 			ID: rs.ID.String(),
// 		})
// 	}
// 	return responses
// }

// List implements service.ConfigService
func (service *configServiceImpl) List(ctx context.Context, offset int, limit int, order string, searchRequest model.ConfigRequest) (responses []model.ConfigResponse) {
	searchEntity := entity.Config{
		ConfigCode: searchRequest.ConfigCode,
	}
	results := service.ConfigRepository.List(ctx, offset, limit, order, searchEntity)
	if len(results) == 0 {
		return []model.ConfigResponse{}
	}

	for _, rs := range results {
		responses = append(responses, model.ConfigResponse{
			ID:          rs.ID.String(),
			ConfigCode:  rs.ConfigCode,
			ConfigValue: rs.ConfigValue,
		})
	}
	return responses
}
