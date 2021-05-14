package main

import (
	"Tarea3/customers"
	"Tarea3/employees"
	"Tarea3/products"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func inicio(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func main() {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	customers.Routes(e)
	employees.Routes(e)
	products.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))

}
