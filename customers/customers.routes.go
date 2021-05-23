package customers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo) {

	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("pia")))
	r.GET("/customers", getCustomer)
	r.POST("/customers", addCustomer)
	r.PUT("/customers", putCustomer)
	r.DELETE("/customers", removeCustomer)
}
