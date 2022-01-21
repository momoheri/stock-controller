package product

type (
	GetQty struct {
		ProductName string `json:"productname" bson:"productname"`
		Quantity    int32  `json:"quantity" bson:"quantity"`
	}

	Validate struct {
		ProductName   string `json:"productname" bson:"productname"`
		ValidateStock int32  `json:"validatestock" bson:"validatestock"`
	}
)
