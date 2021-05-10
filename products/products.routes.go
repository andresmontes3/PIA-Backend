package products

import (
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	e.GET("/products", getProduct)
	e.PUT("/products", putProduct)
	e.DELETE("/products", removeProduct)
	e.POST("/products", addProduct)
}
