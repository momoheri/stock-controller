package product

type (
	Product struct {
		ProductId     string        `json:"id"`
		ProductName   string        `json:"product_name"`
		Category      string        `json:"category"`
		Images        []ImageList   `json:"images"`
		Description   string        `json:"description"`
		Price         int           `json:"price"`
		MinimumOrder  int           `json:"minimum_order"`
		Weight        WeightProduct `json:"weight"`
		Dimensions    Dimensions    `json:"dimensions"`
		StatusProduct int           `json:"status_product"`
		FlashSale     int           `json:"flash_sale"`
	}

	ImageList struct {
		ImageName string `json:"image_name"`
		Path      string `json:"path"`
	}

	WeightProduct struct {
		Unit  string `json:"unit"`
		Value int    `json:"value"`
	}

	Dimensions struct {
		Height int32 `json:"height"`
		Width  int32 `json:"width"`
	}
)
