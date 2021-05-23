package products

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo) {

	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("pia")))
	r.GET("/products", getProduct)
	r.PUT("/products", putProduct)
	r.DELETE("/products", removeProduct)
	r.POST("/products", addProduct)
}
