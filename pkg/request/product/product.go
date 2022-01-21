package product

type (
	ImageList struct {
		ImageName string `json:"image_name" bson:"image_name"`
		Path      string `json:"path" bson:"path"`
	}

	FindeOne struct {
		ProductId string `json:"productid"`
	}
)
