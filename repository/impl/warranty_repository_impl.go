package impl

import (
	"context"
	"fmt"
	"warrantyapi/entity"
	"warrantyapi/exception"
	"warrantyapi/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewWarrantyRepositoryImpl(DB *gorm.DB) repository.WarrantyRepository {
	return &warrantyRepositoryImpl{DB: DB}
}

type warrantyRepositoryImpl struct {
	*gorm.DB
}

// Insert implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) Insert(ctx context.Context, warrantys []entity.Warranty) []entity.Warranty {
	err := repository.DB.WithContext(ctx).Create(&warrantys).Error
	exception.PanicLogging(err)
	return warrantys
}

// GetById implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) GetById(ctx context.Context, id string) entity.Warranty {
	result := entity.Warranty{
		ID: uuid.MustParse(id),
	}
	repository.DB.WithContext(ctx).Debug().
		First(&result)
	return result
}

// Update implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) Update(ctx context.Context, warranty []entity.Warranty) []entity.Warranty {
	err := repository.DB.WithContext(ctx).Save(&warranty).Error
	exception.PanicLogging(err)
	return warranty
}

// Delete implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) Delete(ctx context.Context, warranty entity.Warranty) bool {
	repository.DB.WithContext(ctx).Debug().Delete(&warranty)
	return true
}

// List implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) List(ctx context.Context, offset int, limit int, order string, search entity.Warranty) []entity.Warranty {
	var result []entity.Warranty
	repository.DB.WithContext(ctx).Debug().
		Offset(offset).
		Limit(limit).
		Order(order).
		Where(search).
		Find(&result)
	return result
}

// Total implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) Total(ctx context.Context, search entity.Warranty) int64 {
	var count int64
	repository.DB.WithContext(ctx).Debug().
		Model(&entity.Warranty{}).
		Where(search).
		Count(&count)
	return count
}

// CheckDuplicate implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) CheckDuplicate(ctx context.Context) bool {
	panic("unimplemented")
}

// ListCustomer implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) ListCustomer(ctx context.Context, offset int, limit int, order string, search entity.Warranty) []entity.Warranty {
	// Where("ect_name LIKE ?", "%"+name+"%").
	tx := repository.DB.WithContext(ctx).Debug()
	if search.CustomerEmail != "" {
		tx = tx.Where("customer_email like ?", "%"+search.CustomerEmail+"%")
	}
	if search.CustomerLicensePlate != "" {
		tx = tx.Where("customer_license_plate like ?", "%"+search.CustomerLicensePlate+"%")
	}
	if search.CustomerPhone != "" {
		tx = tx.Where("customer_phone like ?", "%"+search.CustomerPhone+"%")
	}
	var result []entity.Warranty
	tx.Offset(offset).
		Limit(limit).
		Order(order).
		Find(&result)
	return result
}

// Excels implements repository.WarrantyRepository
func (repository *warrantyRepositoryImpl) ListExcels(ctx context.Context, offset int, limit int, order string, search entity.Warranty) []entity.Excels {
	// var result []entity.Warranty
	var (
		result []entity.Excels
		// warrantyTableName = "wt_warranty"
		// productTableName  = "wt_product"
	)
	// repository.DB.WithContext(ctx).Debug().Table("wt_warranty").
	// 	Joins("left join wt_product on wt_warranty.warranty_no = wt_product.warranty_no").
	// 	//  Select("changelog.objectType, changelog.object_type, changelog.object_id, changelog.parent_type, changelog.parent_id, changelog.action, changelog.field, changelog.old_value, changelog.new_value, cc.comment, changelog.created_on, changelog.created_by").
	// 	Select(
	// 		warrantyTableName + ".customer_name, " +
	// 			warrantyTableName + ".warranty_date," +
	// 			warrantyTableName + ".customer_phone, " +
	// 			warrantyTableName + ".customer_license_plate, " +
	// 			warrantyTableName + ".customer_email, " +
	// 			warrantyTableName + ".warranty_no, " +
	// 			warrantyTableName + ".dealer_name, " +
	// 			productTableName + ".product_type, " +
	// 			productTableName + ".product_brand, " +
	// 			productTableName + ".product_amount, " +
	// 			productTableName + ".product_structure_expire, " +
	// 			productTableName + ".product_color_expire, " +
	// 			productTableName + ".product_tire_expire, " +
	// 			productTableName + ".product_mile_expire, " +
	// 			productTableName + ".campagne",
	// 	).
	// 	Find(&result)
	repository.DB.WithContext(ctx).Debug().
		// Offset(offset).
		// Limit(limit).
		// Order(order).
		// Where(search).
		Find(&result)
	// repository.DB.WithContext(ctx).Debug().
	// 	Model(&entity.Warranty{}).
	// 	Select(
	// 		warrantyTableName + ".customer_name, " +
	// 			warrantyTableName + ".warranty_date," +
	// 			warrantyTableName + ".customer_phone, " +
	// 			warrantyTableName + ".customer_license_plate, " +
	// 			warrantyTableName + ".customer_email, " +
	// 			warrantyTableName + ".warranty_no, " +
	// 			warrantyTableName + ".dealer_name, " +
	// 			productTableName + ".product_type, " +
	// 			productTableName + ".product_brand, " +
	// 			productTableName + ".product_amount, " +
	// 			productTableName + ".product_structure_expire, " +
	// 			productTableName + ".product_color_expire, " +
	// 			productTableName + ".product_tire_expire, " +
	// 			productTableName + ".product_mile_expire, " +
	// 			productTableName + ".campagne",
	// 	).
	// 	Joins("left join wt_product on wt_warranty.warranty_no = wt_product.warranty_no").
	// 	Find(&result)
	// Scan(&result)

	fmt.Println(result)
	return result
}
