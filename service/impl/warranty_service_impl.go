package impl

import (
	"context"
	"strconv"
	"warrantyapi/constant"
	"warrantyapi/entity"
	"warrantyapi/model"
	"warrantyapi/repository"
	"warrantyapi/service"

	"github.com/google/uuid"
)

func NewWarrantyServiceImpl(
	warrantyRepository *repository.WarrantyRepository,
	logRepository *repository.LogRepository,
) service.WarrantyService {
	return &warrantyServiceImpl{
		WarrantyRepository: *warrantyRepository,
		LogRepository:      *logRepository,
	}
}

type warrantyServiceImpl struct {
	repository.WarrantyRepository
	repository.LogRepository
}

// Create implements service.WarrantyService
func (service *warrantyServiceImpl) Create(ctx context.Context, warrantyInput []model.WarrantyRequest, createdBy string) []model.WarrantyResponse {
	var warrantys []entity.Warranty
	for _, warranty := range warrantyInput {
		warrantys = append(warrantys, entity.Warranty{
			CreatedBy:  createdBy,
			WarrantyNo: warranty.WarrantyNo,
		})
	}
	service.WarrantyRepository.Insert(ctx, warrantys)
	var responses []model.WarrantyResponse
	for _, rs := range warrantys {
		service.LogRepository.Insert(ctx, entity.Log{
			CreatedBy: createdBy,
			Module:    constant.ModuleWarranty,
			Detail:    "สร้าง : การรับประกัน รหัส  " + rs.ID.String() + " " + rs.WarrantyNo,
		})

		responses = append(responses, model.WarrantyResponse{
			ID: rs.ID.String(),
			// CreatedBy: rs.CreatedBy,
		})
	}
	return responses
}

// FindById implements service.WarrantyService
func (service *warrantyServiceImpl) FindById(ctx context.Context, id int) (responses []model.WarrantyResponse) {
	rs := service.WarrantyRepository.GetById(ctx, id)
	responses = append(responses, model.WarrantyResponse{
		ID: rs.ID.String(),
	})
	return responses
}

// Update implements service.WarrantyService
func (service *warrantyServiceImpl) Update(ctx context.Context, warrantyInput []model.WarrantyRequest, updatedBy string) []model.WarrantyResponse {
	var warrantys []entity.Warranty
	for _, warranty := range warrantyInput {
		warrantys = append(warrantys, entity.Warranty{
			ID: uuid.MustParse(warranty.ID),
		})
	}
	service.WarrantyRepository.Update(ctx, warrantys)
	var responses []model.WarrantyResponse
	for _, rs := range warrantys {
		service.LogRepository.Insert(ctx, entity.Log{
			CreatedBy: updatedBy,
			Module:    constant.ModuleWarranty,
			Detail:    "แก้ไข : การรับประกัน รหัส  " + rs.ID.String() + " " + rs.WarrantyNo,
		})

		responses = append(responses, model.WarrantyResponse{
			ID: rs.ID.String(),
		})
	}
	return responses
}

// Delete implements service.WarrantyService
func (service *warrantyServiceImpl) Delete(ctx context.Context, id int, deletedBy string) bool {
	entityDelete := service.WarrantyRepository.GetById(ctx, id)

	service.WarrantyRepository.Delete(ctx, entityDelete)
	service.LogRepository.Insert(ctx, entity.Log{
		CreatedBy: deletedBy,
		Module:    constant.ModuleWarranty,
		Detail:    "ลบ : การรับประกัน รหัส  " + strconv.FormatUint(uint64(id), 10),
	})

	return true
}

// Create implements service.WarrantyService
func (service *warrantyServiceImpl) List(ctx context.Context, offset int, limit int, order string, warrantyInput model.WarrantyRequest) (responses []model.WarrantyResponse) {

	searchInput := entity.Warranty{
		// Province: warrantyInput.Province,
	}
	results := service.WarrantyRepository.List(ctx, offset, limit, order, searchInput)
	if len(results) == 0 {
		return []model.WarrantyResponse{}
	}

	for _, rs := range results {
		responses = append(responses, model.WarrantyResponse{
			ID: rs.ID.String(),
		})
	}
	return responses
}
