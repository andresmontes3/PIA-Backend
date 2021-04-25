package employees

import (
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	e.GET("/employees", getEmployee)
	e.PUT("/employees", putEmployee)
}
