package product

type StatusPromo int32

const (
	FlashSaleTrue  StatusPromo = 1
	FlashSaleFlase StatusPromo = 0
)

type StatusProduct int32

const (
	Active   StatusProduct = 1
	Inactive StatusProduct = 0
)

type ValidateStock int32

const (
	ValidateTrue  ValidateStock = 1
	ValidateFalse ValidateStock = 0
)
