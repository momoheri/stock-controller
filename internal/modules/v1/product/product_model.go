package product

import (
	"context"
	"log"
	db "stock-controller/internal/db/mongo"
)

func Save(data interface{}) (err error) {
	coll := db.Mongo.Conn.Collection("product")
	_, err = coll.InsertOne(context.TODO(), data)
	if err != nil {
		log.Println("Error!", err)
		return
	}

	return
}
