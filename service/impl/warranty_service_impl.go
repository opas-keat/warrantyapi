package impl

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"warrantyapi/constant"
	"warrantyapi/entity"
	"warrantyapi/model"
	"warrantyapi/repository"
	"warrantyapi/service"

	"github.com/google/uuid"
)

func NewWarrantyServiceImpl(
	warrantyRepository *repository.WarrantyRepository,
	productRepository *repository.ProductRepository,
	logRepository *repository.LogRepository,
) service.WarrantyService {
	return &warrantyServiceImpl{
		WarrantyRepository: *warrantyRepository,
		ProductRepository:  *productRepository,
		LogRepository:      *logRepository,
	}
}

type warrantyServiceImpl struct {
	repository.WarrantyRepository
	repository.ProductRepository
	repository.LogRepository
}

// Create implements service.WarrantyService
func (service *warrantyServiceImpl) Create(ctx context.Context, warrantyInput model.WarrantyRequest, createdBy string) model.WarrantyResponse {
	rand.Seed(time.Now().UnixNano())
	currentTime := time.Now()
	warrantyNo := "WT-" + currentTime.Format("20060102") + currentTime.Format("150405") + strconv.Itoa(rand.Intn(10))
	fmt.Println("warrantyNo = " + warrantyNo)
	var warrantys []entity.Warranty
	warrantys = append(warrantys, entity.Warranty{
		CreatedBy:            createdBy,
		WarrantyNo:           warrantyNo,
		WarrantyDateTime:     time.Now().Format(constant.FORMAT_DATE_TIME),
		DealerCode:           warrantyInput.DealerCode,
		DealerName:           warrantyInput.DealerName,
		CustomerName:         warrantyInput.CustomerName,
		CustomerPhone:        warrantyInput.CustomerPhone,
		CustomerLicensePlate: warrantyInput.CustomerLicensePlate,
		CustomerEmail:        warrantyInput.CustomerEmail,
		CustomerMile:         warrantyInput.CustomerMile,
	})
	service.WarrantyRepository.Insert(ctx, warrantys)
	var responses model.WarrantyResponse
	for _, rs := range warrantys {
		service.LogRepository.Insert(ctx, entity.Log{
			CreatedBy: createdBy,
			Module:    constant.ModuleWarranty,
			Detail:    "สร้าง : การรับประกัน รหัส  " + rs.ID.String() + " " + rs.WarrantyNo,
		})

		responses = model.WarrantyResponse{
			ID:                   rs.ID.String(),
			WarrantyNo:           rs.WarrantyNo,
			WarrantyDateTime:     rs.WarrantyDateTime,
			DealerCode:           rs.DealerCode,
			DealerName:           rs.DealerName,
			CustomerName:         rs.CustomerName,
			CustomerPhone:        rs.CustomerPhone,
			CustomerLicensePlate: rs.CustomerLicensePlate,
			CustomerEmail:        rs.CustomerEmail,
			CustomerMile:         rs.CustomerMile,
		}

		var products []entity.Product
		for _, product := range warrantyInput.ProductRequest {
			products = append(products, entity.Product{
				CreatedBy:     createdBy,
				ProductType:   product.ProductType,
				ProductBrand:  product.ProductBrand,
				ProductAmount: product.ProductAmount,
				WarrantyNo:    rs.WarrantyNo,
			})
		}
		service.ProductRepository.Insert(ctx, products)
		for _, responseProduct := range products {
			responses.ProductResponse = append(responses.ProductResponse, model.ProductResponse{
				ID:            responses.ID,
				ProductType:   responseProduct.ProductType,
				ProductBrand:  responseProduct.ProductBrand,
				ProductAmount: responseProduct.ProductAmount,
				WarrantyNo:    responseProduct.WarrantyNo,
			})
		}
	}
	return responses
}

// FindById implements service.WarrantyService
func (service *warrantyServiceImpl) FindById(ctx context.Context, id int) (responses []model.WarrantyResponse) {
	rs := service.WarrantyRepository.GetById(ctx, id)
	responses = append(responses, model.WarrantyResponse{
		ID:                   rs.ID.String(),
		WarrantyNo:           rs.WarrantyNo,
		WarrantyDateTime:     rs.WarrantyDateTime,
		DealerCode:           rs.DealerCode,
		DealerName:           rs.DealerName,
		CustomerName:         rs.CustomerName,
		CustomerPhone:        rs.CustomerPhone,
		CustomerLicensePlate: rs.CustomerLicensePlate,
		CustomerEmail:        rs.CustomerEmail,
		CustomerMile:         rs.CustomerMile,
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
func (service *warrantyServiceImpl) List(ctx context.Context, offset int, limit int, order string, warrantyInput model.WarrantyRequest) []model.WarrantyResponse {
	searchInput := entity.Warranty{
		CustomerPhone:        warrantyInput.CustomerPhone,
		CustomerLicensePlate: warrantyInput.CustomerLicensePlate,
		CustomerEmail:        warrantyInput.CustomerEmail,
	}
	warrantys := service.WarrantyRepository.List(ctx, offset, limit, order, searchInput)
	if len(warrantys) == 0 {
		return []model.WarrantyResponse{}
	}

	var warrantyResponse []model.WarrantyResponse
	for _, warranty := range warrantys {

		searchProduct := entity.Product{
			WarrantyNo: warranty.WarrantyNo,
		}
		products := service.ProductRepository.List(ctx, offset, limit, "created_at desc", searchProduct)
		var productResponse []model.ProductResponse
		for _, product := range products {
			productResponse = append(productResponse, model.ProductResponse{
				ID:            product.ID.String(),
				ProductType:   product.ProductType,
				ProductBrand:  product.ProductBrand,
				ProductAmount: product.ProductAmount,
				WarrantyNo:    product.WarrantyNo,
			})
		}

		warrantyResponse = append(warrantyResponse, model.WarrantyResponse{
			ID:                   warranty.ID.String(),
			WarrantyNo:           warranty.WarrantyNo,
			WarrantyDateTime:     warranty.WarrantyDateTime,
			DealerCode:           warranty.DealerCode,
			DealerName:           warranty.DealerName,
			CustomerName:         warranty.CustomerName,
			CustomerPhone:        warranty.CustomerPhone,
			CustomerLicensePlate: warranty.CustomerLicensePlate,
			CustomerEmail:        warranty.CustomerEmail,
			CustomerMile:         warranty.CustomerMile,
			ProductResponse:      productResponse,
		})
	}
	return warrantyResponse
}
