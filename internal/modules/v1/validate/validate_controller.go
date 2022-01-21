package validate

import (
	"log"
	"net/http"
	"stock-controller/internal/helper"
	model "stock-controller/pkg/model/product"
	creq "stock-controller/pkg/request/product"
	"stock-controller/pkg/response"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckStock(c echo.Context) error {
	var err error
	id := c.QueryParam("id")

	req := new(creq.FindeOne)
	if err = c.Bind(req); err != nil {
		return response.WriteError(c, response.Body{
			HTTPStatusCode: http.StatusBadRequest,
			Message:        http.StatusText(http.StatusBadRequest),
			Errors:         err,
		})
	}

	req.ProductId = id

	res, err := FindOne(c, req)
	if err != nil {
		statusCode, statusMsg := helper.HTTPCode().MapError(err)
		return response.WriteError(c, response.Body{
			HTTPStatusCode: statusCode,
			Message:        statusMsg,
			Errors:         err,
		})
	}

	return response.WriteSuccess(c, response.Body{
		Status: "Success",
		Data:   res,
	})
}

func ValidateStock(c echo.Context) error {
	var err error
	id := c.QueryParam("id")
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return response.WriteError(c, response.Body{
			HTTPStatusCode: http.StatusBadRequest,
			Message:        http.StatusText(http.StatusBadRequest),
			Errors:         err,
		})
	}
	req := new(model.Product)

	req.Id = &idObj

	if err = c.Bind(req); err != nil {
		return response.WriteError(c, response.Body{
			HTTPStatusCode: http.StatusBadRequest,
			Message:        http.StatusText(http.StatusBadRequest),
			Errors:         err,
		})
	}

	res, err := Update(c, req)
	if err != nil {
		statusCode, statusMsg := helper.HTTPCode().MapError(err)
		return response.WriteError(c, response.Body{
			HTTPStatusCode: statusCode,
			Message:        statusMsg,
			Errors:         err,
		})
	}

	return response.WriteSuccess(c, response.Body{
		Message: "Success update data",
		Data:    res,
	})
}
