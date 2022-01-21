package product

import (
	"net/http"
	model "stock-controller/pkg/model/product"
	"stock-controller/pkg/response"

	"github.com/labstack/echo/v4"
)

func AddProduct(c echo.Context) error {
	var err error
	req := new(model.Product)
	if err = c.Bind(req); err != nil {
		return response.WriteError(c, response.Body{
			HTTPStatusCode: http.StatusBadRequest,
			Message:        http.StatusText(http.StatusBadRequest),
			Errors:         err,
		})
	}

	if err = Save(req); err != nil {
		return response.WriteError(c, response.Body{
			HTTPStatusCode: http.StatusBadRequest,
			Message:        "Error save data",
			Errors:         err,
		})
	}

	return response.WriteSuccess(c, response.Body{
		Status:  "Success",
		Message: "Success save Data",
	})
}
