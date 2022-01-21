package product

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Product struct {
		Id            *primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
		ProductName   string              `json:"productname"`
		Category      string              `json:"category"`
		Images        []ImageList         `json:"images"`
		Description   string              `json:"description"`
		Price         int32               `json:"price"`
		Quantity      int32               `json:"quantity"`
		MinimumOrder  int32               `json:"minimumorder"`
		Weight        WeightProduct       `json:"weight"`
		Dimensions    Dimensions          `json:"dimensions"`
		StatusProduct int32               `json:"statusproduct"`
		ValidateStock int32               `json:"validatestock"`
		FlashSale     int32               `json:"flashsale"`
	}

	ImageList struct {
		Path string `json:"path"`
	}

	WeightProduct struct {
		Unit  string `json:"unit"`
		Value int32  `json:"value"`
	}

	Dimensions struct {
		Height int32 `json:"height"`
		Width  int32 `json:"width"`
	}
)
