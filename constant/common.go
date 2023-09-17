package constant

const (
	ModuleDealer   = "dealer"
	ModuleProduct  = "product"
	ModuleWarranty = "warranty"
)

const (
	STATUS_CODE_OK            = "000"
	STATUS_CODE_BAD_REQUEST   = "400"
	STATUS_CODE_UN_AUTHORIZED = "401"
)

const (
	MESSAGE_LOGIN_FAIL = "ชื่อผู้ใช้งานหรือรหัสผ่าน ไม่ถูกต้อง"

	MESSAGE_STATION_NOT_FOUND = "Station Not Found"

	MESSAGE_USER_NOT_FOUND = "User Not Found"
	MESSAGE_USER_DUPLICATE = "User Duplicated"

	MESSAGE_SUCCESS       = "SUCCESS"
	MESSAGE_BAD_REQUEST   = "BAD REQUEST"
	MESSAGE_UN_AUTHORIZED = "UNAUTHORIZED"
)

const (
	FORMAT_DATE        = "02/01/2006"
	FORMAT_DATE_TIME   = "02/01/2006 15:04:05"
	WarrantyWheelYear  = 6
	WarrantyWheelColor = 6

	WarrantyTireYear = 2
	WarrantyTireMile = 50000

	WarrantyTireYearZestino = 1
	WarrantyTireMileZestino = 40000

	WarrantyPromotionTire = 60
)
