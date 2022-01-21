package product

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	ImageList struct {
		ImageName string `json:"image_name" bson:"image_name"`
		Path      string `json:"path" bson:"path"`
	}

	FindeOne struct {
		Id primitive.ObjectID `json:"id" bson:"_id"`
	}
)
