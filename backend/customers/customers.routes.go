package customers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo) {

	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("pia"))) //se restringen las rutas con la palabra pia

	r.GET("/customers", getCustomer)
	r.POST("/customers", addCustomer)
	r.PUT("/customers", putCustomer)
	r.DELETE("/customers", removeCustomer)

	/*
		r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningMethod: "HS256",
			SigningKey:    []byte("pia"),
			TokenLookup:   "cookie:token",
		}))*/
}
