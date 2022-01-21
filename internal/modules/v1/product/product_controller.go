package product

import (
	"encoding/json"
	"log"
	"net/http"
	model "stock-controller/pkg/model/product"
	"stock-controller/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddProduct(c echo.Context) error {
	var images []model.ImageList
	var weight model.WeightProduct
	var dimensions model.Dimensions

	_, err := c.MultipartForm()
	if err != nil {
		return response.WriteError(c, response.Body{
			HTTPStatusCode: http.StatusBadGateway,
			Message:        http.StatusText(http.StatusInternalServerError),
			Errors:         err,
		})
	}

	price, errPrice := strconv.Atoi(c.FormValue("price"))
	if errPrice != nil {
		log.Println("error!", errPrice.Error())
	}
	min_order, errorder := strconv.Atoi(c.FormValue("minimum_order"))
	if errorder != nil {
		log.Println("error!", errorder.Error())
	}

	if err := json.Unmarshal([]byte(c.FormValue("weight")), &weight); err != nil {
		log.Println("error", err)
	}

	if err := json.Unmarshal([]byte(c.FormValue("images")), &images); err != nil {
		log.Println("error", err)
	}

	if err := json.Unmarshal([]byte(c.FormValue("dimensions")), &dimensions); err != nil {
		log.Println("error", err)
	}

	status, errStatus := strconv.Atoi(c.FormValue("status_product"))
	if errStatus != nil {
		log.Println("error!", errStatus.Error())
	}

	addData := model.Product{
		ProductName:   c.FormValue("product_name"),
		Category:      c.FormValue("category"),
		Images:        images,
		Description:   c.FormValue("description"),
		Price:         price,
		MinimumOrder:  min_order,
		Weight:        weight,
		Dimensions:    dimensions,
		StatusProduct: status,
	}

	err = Save(addData)
	if err != nil {
		return response.WriteError(c, response.Body{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	return response.WriteSuccess(c, response.Body{
		Status:  "Success",
		Message: "Success Save Data",
	})
}
