package validate

import (
	"log"
	db "stock-controller/internal/db/mongo"
	"stock-controller/internal/helper"
	"stock-controller/pkg/model/product"
	creq "stock-controller/pkg/request/product"
	model "stock-controller/pkg/response/product"
	prodtype "stock-controller/pkg/types"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne(c echo.Context, req *creq.FindeOne) (res *model.GetQty, err error) {
	cto, cancel := helper.ContextTimeout(0)
	defer cancel()

	filter := bson.M{
		"_id": req.Id,
	}

	coll := db.Mongo.Conn.Collection("product").FindOne(cto, filter)
	if coll.Err() != nil {
		err = coll.Err()
		c.Logger().Errorf("Error find one. %v", coll.Err())
		return
	}

	if err = coll.Decode(&res); err != nil {
		c.Logger().Errorf("Error decode mongo result, %v", err)
		return
	}

	return
}

func Update(c echo.Context, req *product.Product) (res *model.Validate, err error) {
	cto, cancel := helper.ContextTimeout(0)
	defer cancel()

	filter := bson.M{
		"_id": req.Id,
	}

	update := bson.M{
		"$set": bson.M{"validatestock": prodtype.ValidateTrue},
	}
	log.Println(req.Id)
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	coll := db.Mongo.Conn.Collection("product").FindOneAndUpdate(cto, filter, update, &opts)
	if coll.Err() != nil {
		c.Logger().Errorf("Error Update. %v", coll.Err())
		err = coll.Err()
		return
	}

	if err = coll.Decode(&res); err != nil {
		c.Logger().Errorf("error at decode result. %v", err)
	}

	return
}
