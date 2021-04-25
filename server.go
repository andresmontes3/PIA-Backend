package main

import (
	"Tarea3/customers"
	"Tarea3/employees"
	"Tarea3/products"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/labstack/echo/v4"
)

func inicio(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func main() {
	e := echo.New()
	customers.Routes(e)
	employees.Routes(e)
	products.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
