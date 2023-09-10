package impl

import (
	"context"
	"warrantyapi/entity"
	"warrantyapi/model"
	"warrantyapi/repository"
	"warrantyapi/service"

	"github.com/google/uuid"
)

func NewDealerServiceImpl(dealerRepository *repository.DealerRepository, logRepository *repository.LogRepository) service.DealerService {
	return &dealerServiceImpl{
		DealerRepository: *dealerRepository,
		LogRepository:    *logRepository,
	}
}

type dealerServiceImpl struct {
	repository.DealerRepository
	repository.LogRepository
}

// Create implements service.DealerService
func (service *dealerServiceImpl) Create(ctx context.Context, dealerInput []model.DealerRequest, createdBy string) []model.DealerResponse {
	var dealers []entity.Dealer
	for _, incident := range dealerInput {
		dealers = append(dealers, entity.Dealer{
			CreatedBy:     createdBy,
			DealerCode:    incident.DealerCode,
			DealerName:    incident.DealerName,
			DealerAddress: incident.DealerAddress,
			DealerPhone:   incident.DealerPhone,
			DealerTax:     incident.DealerTax,
			DealerArea:    incident.DealerArea,
		})
	}
	service.DealerRepository.Insert(ctx, dealers)
	var responses []model.DealerResponse
	for _, rs := range dealers {
		service.LogRepository.Insert(ctx, entity.Log{
			CreatedBy: createdBy,
			Module:    "dealer",
			Detail:    "สร้าง : ร้านค้า รหัส  " + rs.ID.String() + " " + rs.DealerCode,
		})

		responses = append(responses, model.DealerResponse{
			ID:            rs.ID.String(),
			DealerCode:    rs.DealerCode,
			DealerName:    rs.DealerName,
			DealerAddress: rs.DealerAddress,
			DealerPhone:   rs.DealerPhone,
			DealerTax:     rs.DealerTax,
			DealerArea:    rs.DealerArea,
		})
	}
	return responses
}

// FindById implements service.DealerService
func (service *dealerServiceImpl) FindById(ctx context.Context, id string) (responses []model.DealerResponse) {
	rs := service.DealerRepository.GetById(ctx, id)
	responses = append(responses, model.DealerResponse{
		ID:            rs.ID.String(),
		DealerCode:    rs.DealerCode,
		DealerName:    rs.DealerName,
		DealerAddress: rs.DealerAddress,
		DealerPhone:   rs.DealerPhone,
		DealerTax:     rs.DealerTax,
		DealerArea:    rs.DealerArea,
	})
	return responses
}

// List implements service.DealerService
func (service *dealerServiceImpl) List(ctx context.Context, offset int, limit int, order string, searchRequest model.DealerRequest) (responses []model.DealerResponse) {
	id, _ := uuid.Parse(searchRequest.ID)
	searchEntity := entity.Dealer{
		ID:         id,
		DealerCode: searchRequest.DealerCode,
	}
	results := service.DealerRepository.List(ctx, offset, limit, order, searchEntity)
	if len(results) == 0 {
		return []model.DealerResponse{}
	}

	for _, rs := range results {
		responses = append(responses, model.DealerResponse{
			ID:            rs.ID.String(),
			DealerCode:    rs.DealerCode,
			DealerName:    rs.DealerName,
			DealerAddress: rs.DealerAddress,
			DealerPhone:   rs.DealerPhone,
			DealerTax:     rs.DealerTax,
			DealerArea:    rs.DealerArea,
		})
	}
	return responses
}
