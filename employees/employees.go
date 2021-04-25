package employees

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/labstack/echo/v4"
)

func getEmployee(c echo.Context) error {

	EmployeeID := c.QueryParam("employeeID")

	employee, err := selectEmployee(EmployeeID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, employee)
}

func selectEmployee(EmployeeID string) (Employee, error) {

	log.Printf("CONECTANDO...\n")
	connectionSql()
	log.Printf("TERMINADO DE CONECTAR\n")

	Employee := Employee{}

	tsql := fmt.Sprintf("SELECT * FROM Employees WHERE EmployeeID = %s;", EmployeeID)
	log.Printf(EmployeeID)
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error al leer las filas " + err.Error())
		return Employee, err
	}
	defer rows.Close()

	for rows.Next() {
		er := rows.Scan(&Employee.EmployeeID, &Employee.LastName, &Employee.FirstName, &Employee.Title,
			&Employee.TitleOfCourtesy, &Employee.BirthDate, &Employee.HireDate, &Employee.Address,
			&Employee.City, &Employee.Region, &Employee.PostalCode, &Employee.Country, &Employee.HomePhone, &Employee.Extension, &Employee.Photo,
			&Employee.Notes, &Employee.ReportsTo, &Employee.PhotoPath)
		if err != nil {
			fmt.Println("Error al leer las filas: " + err.Error())
			return Employee, er
		}

	}
	defer db.Close()
	return Employee, nil
}

func putEmployee(c echo.Context) error {
	employee := Employee{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&employee)
	if err != nil {
		log.Printf("fallo al procesar addEmployee: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	result, err := updateEmployee(employee.EmployeeID, employee.LastName, employee.FirstName, employee.Title,
		employee.TitleOfCourtesy, employee.BirthDate, employee.HireDate, employee.Address,
		employee.City, employee.Region, employee.PostalCode, employee.Country, employee.HomePhone, employee.Extension, employee.Photo,
		employee.Notes, employee.ReportsTo, employee.PhotoPath)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)

	log.Printf("Este es un cliente: %#v", employee)

	return c.String(http.StatusOK, "tenemos tu cliente")

}

func updateEmployee(EmployeeID int, LastName string, FirstName string, Title string, TitleOfCourtesy string,
	BirthDate string, HireDate string, Address string, City string, Region string, PostalCode string,
	Country string, HomePhone string, Extension string, Photo string, Notes string, ReportsTo int, PhotoPath string) (int64, error) {

	log.Printf("CONECTANDO...\n")
	connectionSql()
	log.Printf("TERMINADO DE CONECTAR\n")

	tsql := fmt.Sprintf("UPDATE Employees SET LastName='%s',FirstName='%s',Title='%s',TitleOfCourtesy='%s',BirthDate='%s',Hiredate ='%s',Address='%s',City='%s',Region='%s',PostalCode='%s',Country ='%s',HomePhone='%s',Extension='%s',Photo='%s',Notes='%s',ReportsTo = %d, PhotoPath='%s' WHERE EmployeeID = %d;",
		LastName, FirstName, Title,
		TitleOfCourtesy, BirthDate, HireDate, Address,
		City, Region, PostalCode, Country, HomePhone, Extension, Photo,
		Notes, ReportsTo, PhotoPath, EmployeeID)

	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error al actualizar la fila " + err.Error())
		return -1, err
	}

	defer db.Close()
	return result.RowsAffected()
}

var server = "localhost"
var port = 1433
var user = ""
var password = ""
var database = "Northwind"
var db *sql.DB

func connectionSql() {
	var err error
	// Create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")
}
