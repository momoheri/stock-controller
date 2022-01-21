package product

type StatusPromo int8

const (
	FlashSaleTrue  StatusPromo = 1
	FlashSaleFlase StatusPromo = 0
)

type StatusProduct int8

const (
	Active   StatusProduct = 1
	Unactive StatusProduct = 0
)
