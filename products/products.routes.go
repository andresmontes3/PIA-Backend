package products

import (
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	e.GET("/products", getProduct)
	e.PUT("/products", putProduct)
}
