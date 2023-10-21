package impl

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"
	_ "time/tzdata"
	"warrantyapi/common"
	"warrantyapi/constant"
	"warrantyapi/entity"
	"warrantyapi/model"
	"warrantyapi/repository"
	"warrantyapi/service"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func NewWarrantyServiceImpl(
	warrantyRepository *repository.WarrantyRepository,
	productRepository *repository.ProductRepository,
	configRepository *repository.ConfigRepository,
	logRepository *repository.LogRepository,
) service.WarrantyService {
	return &warrantyServiceImpl{
		WarrantyRepository: *warrantyRepository,
		ProductRepository:  *productRepository,
		ConfigRepository:   *configRepository,
		LogRepository:      *logRepository,
	}
}

type warrantyServiceImpl struct {
	repository.WarrantyRepository
	repository.ProductRepository
	repository.ConfigRepository
	repository.LogRepository
}

// Create implements service.WarrantyService
func (service *warrantyServiceImpl) Create(ctx context.Context, warrantyInput model.WarrantyRequest, createdBy string) model.WarrantyResponse {
	// rand.Seed(time.Now().UnixNano())
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// currentTime := time.Now()
	// warrantyNo := "WT-" + currentTime.Format("20060102") + currentTime.Format("150405") + strconv.Itoa(1000+r.Intn(10000-1000))
	// warrantyNo := currentTime.Format("0601") + strconv.Itoa(100000+r.Intn(1000000-10000))
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}
	configs := service.ConfigRepository.List(ctx, 0, 100, "", entity.Config{})
	configCal := common.SetConfigCal(configs)
	log.Debug().
		Any("configCal", configCal).
		Send()
	total := service.WarrantyRepository.Total(ctx, entity.Warranty{})
	//2310000002
	warrantyNo := time.Now().Format("0601") + fmt.Sprintf("%06d", total+1)
	fmt.Println("warrantyNo = " + warrantyNo)
	var warrantys []entity.Warranty
	warrantys = append(warrantys, entity.Warranty{
		CreatedBy:            createdBy,
		WarrantyNo:           warrantyNo,
		WarrantyDateTime:     time.Now().In(loc).Format(constant.FORMAT_DATE_TIME),
		DealerCode:           warrantyInput.DealerCode,
		DealerName:           warrantyInput.DealerName,
		CustomerName:         warrantyInput.CustomerName,
		CustomerPhone:        warrantyInput.CustomerPhone,
		CustomerLicensePlate: warrantyInput.CustomerLicensePlate,
		CustomerEmail:        warrantyInput.CustomerEmail,
		CustomerMile:         warrantyInput.CustomerMile,
	})
	log.Debug().
		Any("warrantys", warrantys).
		Send()
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
		//Insert Product
		var products []entity.Product
		for _, product := range warrantyInput.ProductRequest {
			wheelStructureExpire := common.CalculateYearExpire(configCal.WarrantyWheelYear)
			wheelColorExpire := common.CalculateMonthExpire(configCal.WarrantyWheelColor)
			tireExpire := common.CalculateYearExpire(configCal.WarrantyTireYear)
			tireMileExpire := common.CalculateMileExpire(warrantyInput.CustomerMile, configCal.WarrantyTireMile)
			promotionExpire := ""
			if strings.EqualFold(product.ProductBrand, "zestino") {
				tireExpire = common.CalculateYearExpire(configCal.WarrantyTireYearZestino)
				tireMileExpire = common.CalculateMileExpire(warrantyInput.CustomerMile, configCal.WarrantyTireMileZestino)
			}
			if strings.EqualFold(product.ProductType, "tire") && (product.ProductAmount >= 4) {
				promotionExpire = common.CalculateDayExpire(configCal.WarrantyPromotionTire)
			}
			products = append(products, entity.Product{
				CreatedBy:               createdBy,
				ProductType:             product.ProductType,
				ProductBrand:            product.ProductBrand,
				ProductAmount:           product.ProductAmount,
				ProductStructureExpire:  wheelStructureExpire,
				ProductColorExpire:      wheelColorExpire,
				ProductTireExpire:       tireExpire,
				ProductMileExpire:       tireMileExpire,
				ProductPromotionExpire:  promotionExpire,
				WarrantyNo:              warrantyNo,
				WarrantyWheelYear:       configCal.WarrantyWheelYear,
				WarrantyWheelColor:      configCal.WarrantyWheelColor,
				WarrantyTireYear:        configCal.WarrantyTireYear,
				WarrantyTireMile:        configCal.WarrantyTireMile,
				WarrantyTireYearZestino: configCal.WarrantyTireYearZestino,
				WarrantyTireMileZestino: configCal.WarrantyTireMileZestino,
				Promotion:               configCal.Campagne,
				PromotionDay:            configCal.WarrantyPromotionTire,
				PromotionMile:           0,
			})
		}
		log.Debug().
			Any("products", products).
			Send()
		responseProducts := service.ProductRepository.Insert(ctx, products)
		for _, responseProduct := range responseProducts {
			responses.ProductResponse = append(responses.ProductResponse, model.ProductResponse{
				ID:                      responseProduct.ID.String(),
				ProductType:             responseProduct.ProductType,
				ProductBrand:            responseProduct.ProductBrand,
				ProductAmount:           responseProduct.ProductAmount,
				ProductStructureExpire:  responseProduct.ProductStructureExpire,
				ProductColorExpire:      responseProduct.ProductColorExpire,
				ProductTireExpire:       responseProduct.ProductTireExpire,
				ProductMileExpire:       responseProduct.ProductMileExpire,
				ProductPromotionExpire:  responseProduct.ProductPromotionExpire,
				WarrantyNo:              responseProduct.WarrantyNo,
				WarrantyWheelYear:       responseProduct.WarrantyWheelYear,
				WarrantyWheelColor:      responseProduct.WarrantyWheelColor,
				WarrantyTireYear:        responseProduct.WarrantyTireYear,
				WarrantyTireMile:        responseProduct.WarrantyTireMile,
				WarrantyTireYearZestino: responseProduct.WarrantyTireYearZestino,
				WarrantyTireMileZestino: responseProduct.WarrantyTireMileZestino,
				Promotion:               responseProduct.Promotion,
				PromotionDay:            responseProduct.PromotionDay,
				PromotionMile:           responseProduct.PromotionMile,
			})
		}
	}
	return responses
}

// FindById implements service.WarrantyService
func (service *warrantyServiceImpl) FindById(ctx context.Context, id string) model.WarrantyResponse {
	rs := service.WarrantyRepository.GetById(ctx, id)
	// var responses model.WarrantyResponse
	responses := model.WarrantyResponse{
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
	ProductSearch := entity.Product{
		WarrantyNo: rs.WarrantyNo,
	}
	responseProducts := service.ProductRepository.List(ctx, 0, 100, "product_type desc", ProductSearch)
	for _, responseProduct := range responseProducts {
		responses.ProductResponse = append(responses.ProductResponse, model.ProductResponse{
			ID:                      responseProduct.ID.String(),
			ProductType:             responseProduct.ProductType,
			ProductBrand:            responseProduct.ProductBrand,
			ProductAmount:           responseProduct.ProductAmount,
			ProductStructureExpire:  responseProduct.ProductStructureExpire,
			ProductColorExpire:      responseProduct.ProductColorExpire,
			ProductTireExpire:       responseProduct.ProductTireExpire,
			ProductMileExpire:       responseProduct.ProductMileExpire,
			ProductPromotionExpire:  responseProduct.ProductPromotionExpire,
			WarrantyNo:              responseProduct.WarrantyNo,
			WarrantyWheelYear:       responseProduct.WarrantyWheelYear,
			WarrantyWheelColor:      responseProduct.WarrantyWheelColor,
			WarrantyTireYear:        responseProduct.WarrantyTireYear,
			WarrantyTireMile:        responseProduct.WarrantyTireMile,
			WarrantyTireYearZestino: responseProduct.WarrantyTireYearZestino,
			WarrantyTireMileZestino: responseProduct.WarrantyTireMileZestino,
			Promotion:               responseProduct.Promotion,
			PromotionDay:            responseProduct.PromotionDay,
			PromotionMile:           responseProduct.PromotionMile,
		})
	}
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
func (service *warrantyServiceImpl) Delete(ctx context.Context, id string, deletedBy string) bool {
	entityDelete := service.WarrantyRepository.GetById(ctx, id)

	service.WarrantyRepository.Delete(ctx, entityDelete)
	service.LogRepository.Insert(ctx, entity.Log{
		CreatedBy: deletedBy,
		Module:    constant.ModuleWarranty,
		Detail:    "ลบ : การรับประกัน รหัส  " + id,
	})

	return true
}

// ListCustomer implements service.WarrantyService
func (service *warrantyServiceImpl) ListCustomer(ctx context.Context, offset int, limit int, order string, warrantyInput model.WarrantyRequest) []model.WarrantyResponse {
	searchInput := entity.Warranty{
		CustomerPhone:        warrantyInput.CustomerPhone,
		CustomerLicensePlate: warrantyInput.CustomerLicensePlate,
		CustomerEmail:        warrantyInput.CustomerEmail,
	}
	warrantys := service.WarrantyRepository.ListCustomer(ctx, offset, limit, order, searchInput)
	if len(warrantys) == 0 {
		return []model.WarrantyResponse{}
	}

	var warrantyResponse []model.WarrantyResponse
	for _, warranty := range warrantys {
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
		})
	}
	return warrantyResponse
}

// List implements service.WarrantyService
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
				ID:                      product.ID.String(),
				ProductType:             product.ProductType,
				ProductBrand:            product.ProductBrand,
				ProductAmount:           product.ProductAmount,
				ProductStructureExpire:  product.ProductStructureExpire,
				ProductColorExpire:      product.ProductColorExpire,
				ProductTireExpire:       product.ProductTireExpire,
				ProductMileExpire:       product.ProductMileExpire,
				ProductPromotionExpire:  product.ProductPromotionExpire,
				WarrantyNo:              product.WarrantyNo,
				WarrantyWheelYear:       product.WarrantyWheelYear,
				WarrantyWheelColor:      product.WarrantyWheelColor,
				WarrantyTireYear:        product.WarrantyTireYear,
				WarrantyTireMile:        product.WarrantyTireMile,
				WarrantyTireYearZestino: product.WarrantyTireYearZestino,
				WarrantyTireMileZestino: product.WarrantyTireMileZestino,
				Promotion:               product.Promotion,
				PromotionDay:            product.PromotionDay,
				PromotionMile:           product.PromotionMile,
			})
		}
		imageFolder := warranty.WarrantyDateTime[6:10] + warranty.WarrantyDateTime[3:5] + warranty.WarrantyDateTime[0:2]
		pathImageCar := ""
		pathImageReceive := ""
		files, _ := filepath.Glob("uploads/" + imageFolder + "/" + warranty.WarrantyNo + "_car*")
		fmt.Println(len(files))
		for _, file := range files {
			pathImageCar = file
		}
		filesReceive, _ := filepath.Glob("uploads/" + imageFolder + "/" + warranty.WarrantyNo + "_receive*")
		fmt.Println(len(filesReceive))
		for _, file := range filesReceive {
			pathImageReceive = file
		}
		// pathImageCar := "/uploads/" + imageFolder + "/" + warranty.WarrantyNo + "_car.png"
		// pathImageReceive := "/uploads/" + imageFolder + "/" + warranty.WarrantyNo + "_receive.png"
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
			UrlCar:               pathImageCar,
			UrlReceive:           pathImageReceive,
		})
	}
	return warrantyResponse
}
