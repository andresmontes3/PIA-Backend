package customers

import (
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	e.GET("/customers", getCustomer)
	e.POST("/customers", addCustomer)
	e.PUT("/customers", putCustomer)
	e.DELETE("/customers", removeCustomer)
}
