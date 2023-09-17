package common

import (
	"strconv"
	"time"
	"warrantyapi/constant"
)

//	func CalculateWheelExpire() string {
//		// t1 := time.Now()
//		// t2 := t1.AddDate(constant.WarrantyWheelYear, 0, 0)
//		return time.Now().AddDate(constant.WarrantyWheelYear, 0, 0).Format(constant.FORMAT_DATE)
//	}
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
