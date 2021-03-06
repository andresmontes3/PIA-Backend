package employees

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo) {

	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("pia"))) //se restringen las rutas con la palabra pia
	r.GET("/employees", getEmployee)
	r.PUT("/employees", putEmployee)
	r.DELETE("/employees", removeEmployee)
	r.POST("/employees", addEmployee)
}
