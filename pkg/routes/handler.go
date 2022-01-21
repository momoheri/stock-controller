package routes

import (
	prod "stock-controller/internal/modules/v1/product"
	val "stock-controller/internal/modules/v1/validate"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//Post Data
	e.POST("/v1/AddUser", prod.AddProduct)

	//Get Data
	e.GET("/v1/CheckValidate", val.CheckValidate)

	//Update
	e.PUT("/v1/validate", val.ValidateStock)

	return e
}
