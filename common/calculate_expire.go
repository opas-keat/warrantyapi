package common

import (
	"strconv"
	"time"
	"warrantyapi/constant"
	"warrantyapi/entity"

	"github.com/rs/zerolog/log"
)

//	func CalculateWheelExpire() string {
//		// t1 := time.Now()
//		// t2 := t1.AddDate(constant.WarrantyWheelYear, 0, 0)
//		return time.Now().AddDate(constant.WarrantyWheelYear, 0, 0).Format(constant.FORMAT_DATE)
//	}
type ConfigCalculate struct {
	WarrantyWheelYear       int
	WarrantyWheelColor      int
	WarrantyTireYear        int
	WarrantyTireMile        int
	WarrantyTireYearZestino int
	WarrantyTireMileZestino int
	WarrantyPromotionTire   int
	Campagne                string
}

func CalculateYearExpire(year int) string {
	return time.Now().AddDate(year, 0, 0).Format(constant.FORMAT_DATE)
}

func CalculateMonthExpire(month int) string {
	return time.Now().AddDate(0, month, 0).Format(constant.FORMAT_DATE)
}

func CalculateDayExpire(day int) string {
	return time.Now().AddDate(0, 0, day).Format(constant.FORMAT_DATE)
}

func CalculateMileExpire(mile string, milePlus int) string {
	m, _ := strconv.Atoi(mile)
	// mileExpire := m + constant.WarrantyTireMile
	return strconv.Itoa(m + milePlus)
}

func SetConfigCal(configs []entity.Config) ConfigCalculate {
	var configCal ConfigCalculate
	for _, config := range configs {
		log.Debug().
			Str("config_code", config.ConfigCode).
			Str("config_value", config.ConfigValue).
			Caller().
			Send()
		if config.ConfigCode == "WarrantyWheelYear" {
			configCal.WarrantyWheelYear, _ = strconv.Atoi(config.ConfigValue)
		} else if config.ConfigCode == "WarrantyWheelColor" {
			configCal.WarrantyWheelColor, _ = strconv.Atoi(config.ConfigValue)
		} else if config.ConfigCode == "WarrantyTireYear" {
			configCal.WarrantyTireYear, _ = strconv.Atoi(config.ConfigValue)
		} else if config.ConfigCode == "WarrantyTireMile" {
			configCal.WarrantyTireMile, _ = strconv.Atoi(config.ConfigValue)
		} else if config.ConfigCode == "WarrantyTireYearZestino" {
			configCal.WarrantyTireYearZestino, _ = strconv.Atoi(config.ConfigValue)
		} else if config.ConfigCode == "WarrantyTireMileZestino" {
			configCal.WarrantyTireMileZestino, _ = strconv.Atoi(config.ConfigValue)
		} else if config.ConfigCode == "WarrantyPromotionTire" {
			configCal.WarrantyPromotionTire, _ = strconv.Atoi(config.ConfigValue)
		} else if config.ConfigCode == "Campange" {
			configCal.Campagne = config.ConfigValue
		}
	}
	return configCal
}
