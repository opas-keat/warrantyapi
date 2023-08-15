package impl

import (
	"context"
	"warrantyapi/entity"
	"warrantyapi/model"
	"warrantyapi/repository"
	"warrantyapi/service"
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

// Create implements service.Create
func (service *dealerServiceImpl) Create(ctx context.Context, dealerInput model.DealerRequest, createdBy string) model.DealerResponse {
	dealer := entity.Dealer{
		CreatedBy:     createdBy,
		DealerCode:    dealerInput.DealerCode,
		DealerName:    dealerInput.DealerName,
		DealerAddress: dealerInput.DealerAddress,
		DealerPhone:   dealerInput.DealerPhone,
		DealerTax:     dealerInput.DealerTax,
	}
	dealer = service.DealerRepository.Insert(ctx, dealer)
	service.LogRepository.Insert(ctx, entity.Log{
		CreatedBy: createdBy,
		Module:    "dealer",
		// Detail: dealer.ID.String(),
		Detail: "สร้าง : dealer รหัส  " + dealer.ID.String() + " ชื่อ " + dealer.DealerName,
	})
	return model.DealerResponse{
		ID:            dealer.ID.String(),
		DealerCode:    dealer.DealerCode,
		DealerName:    dealer.DealerName,
		DealerAddress: dealer.DealerAddress,
		DealerPhone:   dealer.DealerPhone,
		DealerTax:     dealer.DealerTax,
	}
}
