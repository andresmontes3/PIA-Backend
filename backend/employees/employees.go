package employees

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func dbConnect() *gorm.DB {
	log.Printf("CONECTANDO...\n")

	dsn := "sqlserver://@localhost:1433?database=Backend"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Printf("Connected!\n")
	return db
}

//función que se llama al hacer GET
func getEmployee(c echo.Context) error {

	EmployeeID := c.QueryParam("employee_id")
	log.Printf(EmployeeID)
	employees, err := selectEmployee(EmployeeID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, employees) //regresa un json con los empleados
}

func selectEmployee(EmployeeID string) ([]Employee, error) {
	//se conecta a la base de datos
	db := dbConnect()

	employees := []Employee{}
	//Si se ingresó un ID, se busca ese ID, de lo contrario se muestra la lista completa
	if EmployeeID != "" {
		db.Where("employee_id = ?", EmployeeID).First(&employees)

		return employees, nil
	} else {

		result := db.Order("last_name").Find(&employees)

		return employees, result.Error
	}

}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
//función que se llama al hacer POST
func addEmployee(c echo.Context) error {
	employee := Employee{}

	defer c.Request().Body.Close()
	//obtiene los datos del JSON y los almacena en employee de tipo Employee
	err := json.NewDecoder(c.Request().Body).Decode(&employee)
	if err != nil {
		log.Printf("fallo al procesar addEmployee: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	//Se envian a una función para trabajar con ellos en la base de datos
	result, err := insertEmployee(employee.EmployeeID, employee.LastName, employee.FirstName, employee.Title,
		employee.HireDate, employee.Address, employee.City, employee.Region, employee.PostalCode, employee.Phone)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)

	log.Printf("Este es un empleado: %#v", employee)

	return c.String(http.StatusOK, "tenemos tu empleado")
}

func insertEmployee(employeeID int, lastName string, firstName string, title string, hireDate string,
	address string, city string, region string, postalCode string, phone string) (int64, error) {
	//se conecta a la base de datos
	db := dbConnect()
	//se crea un nuevo employee con los datos recibidos y se insertan con el db.Create
	employee := Employee{EmployeeID: employeeID, LastName: lastName, FirstName: firstName, Title: title,
		HireDate: hireDate, Address: address, City: city, Region: region, PostalCode: postalCode, Phone: phone}
	result := db.Create(&employee)

	return result.RowsAffected, result.Error

}

/////////////////////////////////////////////////////////////////////////
//función llamada al hacer PUT
func putEmployee(c echo.Context) error {
	employee := Employee{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&employee)
	if err != nil {
		log.Printf("fallo al procesar putEmployee: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	//Al igual que en post se reciben todos los parámetros a través de un JSON
	result, err := updateEmployee(employee.EmployeeID, employee.LastName, employee.FirstName, employee.Title,
		employee.HireDate, employee.Address, employee.City, employee.Region, employee.PostalCode, employee.Phone)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)

	log.Printf("Este es un cliente: %#v", employee)

	return c.String(http.StatusOK, "tenemos tu empleado")

}

func updateEmployee(employeeID int, lastName string, firstName string, title string, hireDate string,
	address string, city string, region string, postalCode string, phone string) (int64, error) {
	//se conecta a la base de datos
	db := dbConnect()
	//se hace el UPDATE con un WHERE donde el ID introducido sea el mismo
	result := db.Model(&Employee{}).Where("employee_id = ?", employeeID).Updates(map[string]interface{}{
		"last_name": lastName, "first_name": firstName, "title": title, "hire_date": hireDate,
		"address": address, "city": city, "region": region, "postal_code": postalCode, "phone": phone})

	return result.RowsAffected, result.Error
}

///////////////////////////////////////////////////////////////////////////////////////////////
//función llamada al hacer DELETE
func removeEmployee(c echo.Context) error {
	//Se almacena el ID introducido
	EmployeeID := c.QueryParam("employee_id")

	result, err := deleteEmployee(EmployeeID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)
	return c.String(http.StatusOK, "Empleado eliminado")
}

func deleteEmployee(employeeID string) (int64, error) {
	//se conecta a la base de datos
	db := dbConnect()
	//Donde el ID sea igual, ejecuta un DELETE
	result := db.Where("employee_id = ?", employeeID).Delete(&Employee{})
	return result.RowsAffected, result.Error
}
