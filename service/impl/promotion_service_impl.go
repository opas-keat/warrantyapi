package impl

import (
	"context"
	"strings"
	"time"
	"warrantyapi/constant"
	"warrantyapi/entity"
	"warrantyapi/model"
	"warrantyapi/repository"
	"warrantyapi/service"

	"github.com/google/uuid"
)

func NewPromotionServiceImpl(
	promotionRepository *repository.PromotionRepository,
	logRepository *repository.LogRepository,
) service.PromotionService {
	return &promotionServiceImpl{
		PromotionRepository: *promotionRepository,
		LogRepository:       *logRepository,
	}
}

type promotionServiceImpl struct {
	repository.PromotionRepository
	repository.LogRepository
}

// Create implements service.PromotionService
func (service *promotionServiceImpl) Create(ctx context.Context, promotionInput []model.PromotionRequest, createdBy string) []model.PromotionResponse {
	var promotions []entity.Promotion
	for _, promotion := range promotionInput {
		dateForm, _ := time.Parse(constant.FORMAT_DATE, promotion.PromotionFrom)
		dateTo, _ := time.Parse(constant.FORMAT_DATE, promotion.PromotionTo)
		promotions = append(promotions, entity.Promotion{
			CreatedBy:            createdBy,
			PromotionType:        promotion.PromotionType,
			PromotionBrand:       strings.ToUpper(promotion.PromotionBrand),
			PromotionDetail:      promotion.PromotionDetail,
			PromotionWarrantyDay: promotion.PromotionWarrantyDay,
			PromotionStatus:      promotion.PromotionStatus,
			PromotionFrom:        dateForm,
			PromotionTo:          dateTo,
		})
	}
	service.PromotionRepository.Insert(ctx, promotions)
	var responses []model.PromotionResponse
	for _, rs := range promotions {
		service.LogRepository.Insert(ctx, entity.Log{
			CreatedBy: createdBy,
			Module:    constant.ModulePromotion,
			Detail:    "สร้าง : แคมเปญ รหัส  " + rs.ID.String() + " " + rs.PromotionDetail,
		})

		responses = append(responses, model.PromotionResponse{
			ID:                   rs.ID.String(),
			PromotionType:        rs.PromotionType,
			PromotionBrand:       rs.PromotionBrand,
			PromotionDetail:      rs.PromotionDetail,
			PromotionWarrantyDay: rs.PromotionWarrantyDay,
			PromotionStatus:      rs.PromotionStatus,
			PromotionFrom:        rs.PromotionFrom.Format(constant.FORMAT_DATE),
			PromotionTo:          rs.PromotionTo.Format(constant.FORMAT_DATE),
		})
	}
	return responses
}

// Update implements service.PromotionService
func (service *promotionServiceImpl) Update(ctx context.Context, promotionInput []model.PromotionRequest, updatedBy string) []model.PromotionResponse {
	var promotions []entity.Promotion
	for _, promotion := range promotionInput {
		dateForm, _ := time.Parse(constant.FORMAT_DATE, promotion.PromotionFrom)
		dateTo, _ := time.Parse(constant.FORMAT_DATE, promotion.PromotionTo)
		promotions = append(promotions, entity.Promotion{
			ID:                   uuid.MustParse(promotion.ID),
			PromotionType:        promotion.PromotionType,
			PromotionBrand:       promotion.PromotionBrand,
			PromotionDetail:      promotion.PromotionDetail,
			PromotionWarrantyDay: promotion.PromotionWarrantyDay,
			PromotionStatus:      promotion.PromotionStatus,
			PromotionFrom:        dateForm,
			PromotionTo:          dateTo,
		})
	}
	service.PromotionRepository.Update(ctx, promotions)
	var responses []model.PromotionResponse
	for _, rs := range promotions {
		service.LogRepository.Insert(ctx, entity.Log{
			CreatedBy: updatedBy,
			Module:    constant.ModuleWarranty,
			Detail:    "แก้ไข : แคมเปญ รหัส  " + rs.ID.String() + " " + rs.PromotionDetail,
		})

		responses = append(responses, model.PromotionResponse{
			ID:                   rs.ID.String(),
			PromotionType:        rs.PromotionType,
			PromotionBrand:       rs.PromotionBrand,
			PromotionDetail:      rs.PromotionDetail,
			PromotionWarrantyDay: rs.PromotionWarrantyDay,
			PromotionStatus:      rs.PromotionStatus,
			PromotionFrom:        rs.PromotionFrom.Format(constant.FORMAT_DATE),
			PromotionTo:          rs.PromotionTo.Format(constant.FORMAT_DATE),
		})
	}
	return responses
}

// List implements service.PromotionService
func (service *promotionServiceImpl) List(ctx context.Context, offset int, limit int, order string, searchRequest model.PromotionRequest) (responses []model.PromotionResponse) {
	searchEntity := entity.Promotion{}
	results := service.PromotionRepository.List(ctx, offset, limit, order, searchEntity)
	if len(results) == 0 {
		return []model.PromotionResponse{}
	}

	for _, rs := range results {
		responses = append(responses, model.PromotionResponse{
			ID:                   rs.ID.String(),
			PromotionType:        rs.PromotionType,
			PromotionBrand:       rs.PromotionBrand,
			PromotionDetail:      rs.PromotionDetail,
			PromotionWarrantyDay: rs.PromotionWarrantyDay,
			PromotionStatus:      rs.PromotionStatus,
			PromotionFrom:        rs.PromotionFrom.Format(constant.FORMAT_DATE),
			PromotionTo:          rs.PromotionTo.Format(constant.FORMAT_DATE),
		})
	}
	return responses
}

// ListActivePromotion implements service.PromotionService
func (service *promotionServiceImpl) ListActivePromotion(
	ctx context.Context,
	offset int, limit int, order string,
	searchRequest model.PromotionRequest,
	warrantyCreated time.Time,
) (responses []model.PromotionResponse) {
	searchEntity := entity.Promotion{
		PromotionStatus: searchRequest.PromotionStatus,
		PromotionType:   searchRequest.PromotionType,
		PromotionBrand:  searchRequest.PromotionBrand,
	}
	results := service.PromotionRepository.ListActivePromotion(ctx, offset, limit, order, searchEntity, warrantyCreated)
	if len(results) == 0 {
		return []model.PromotionResponse{}
	}

	for _, rs := range results {
		responses = append(responses, model.PromotionResponse{
			ID:                   rs.ID.String(),
			PromotionType:        rs.PromotionType,
			PromotionBrand:       rs.PromotionBrand,
			PromotionDetail:      rs.PromotionDetail,
			PromotionWarrantyDay: rs.PromotionWarrantyDay,
			PromotionStatus:      rs.PromotionStatus,
			PromotionFrom:        rs.PromotionFrom.Format(constant.FORMAT_DATE),
			PromotionTo:          rs.PromotionTo.Format(constant.FORMAT_DATE),
		})
	}
	return responses
}
