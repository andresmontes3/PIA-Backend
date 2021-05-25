package customers

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
func getCustomer(c echo.Context) error {

	CustomerID := c.QueryParam("customer_id")
	log.Printf(CustomerID)
	customers, err := selectCustomer(CustomerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, customers) //regresa un json con los clientes
}

func selectCustomer(CustomerID string) ([]Customer, error) {

	//se conecta a la base de datos
	db := dbConnect()

	customers := []Customer{}

	//Si se ingresó un ID, se busca ese ID, de lo contrario se muestra la lista completa

	if CustomerID != "" {
		db.Where("customer_id = ?", CustomerID).First(&customers)

		return customers, nil
	} else {

		result := db.Order("company_name").Find(&customers)

		return customers, result.Error
	}

}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

//función que se llama al hacer POST
func addCustomer(c echo.Context) error {

	customer := Customer{}

	defer c.Request().Body.Close()

	//obtiene los datos del JSON y los almacena en customer de tipo Customer
	err := json.NewDecoder(c.Request().Body).Decode(&customer)
	if err != nil {
		log.Printf("fallo al procesar addCustomer: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	//Se envian a una función para trabajar con ellos en la base de datos
	result, err := insertCustomer(customer.CustomerID, customer.CompanyName, customer.ContactName, customer.ContactTitle,
		customer.Address, customer.City, customer.Region, customer.PostalCode,
		customer.Country, customer.Phone)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)

	log.Printf("Este es un cliente: %#v", customer)

	return c.String(http.StatusOK, "tenemos tu cliente")
}

func insertCustomer(customerID string, companyName string, contactName string, contactTitle string, address string, city string,
	region string, postalCode string, country string, phone string) (int64, error) {

	//se conecta a la base de datos
	db := dbConnect()

	//se crea un nuevo customer con los datos recibidos y se insertan con el db.Create
	customer := Customer{CustomerID: customerID, CompanyName: companyName, ContactName: contactName,
		ContactTitle: contactTitle, Address: address, City: city, Region: region, PostalCode: postalCode,
		Country: country, Phone: phone}
	result := db.Create(&customer)

	return result.RowsAffected, result.Error

}

/////////////////////////////////////////////////////////////////////////
//función llamada al hacer PUT
func putCustomer(c echo.Context) error {
	customer := Customer{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&customer)
	if err != nil {
		log.Printf("fallo al procesar addCustomer: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	//Al igual que en post se reciben todos los parámetros a través de un JSON
	result, err := updateCustomer(customer.CustomerID, customer.CompanyName, customer.ContactName, customer.ContactTitle,
		customer.Address, customer.City, customer.Region, customer.PostalCode,
		customer.Country, customer.Phone)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)

	log.Printf("Este es un cliente: %#v", customer)

	return c.String(http.StatusOK, "tenemos tu cliente")

}

func updateCustomer(customerID string, companyName string, contactName string, contactTitle string, address string, city string,
	region string, postalCode string, country string, phone string) (int64, error) {

	//se conecta a la base de datos
	db := dbConnect()

	//se hace el UPDATE con un WHERE donde el ID introducido sea el mismo

	result := db.Model(&Customer{}).Where("customer_id = ?", customerID).Updates(map[string]interface{}{"company_name": companyName,
		"contact_name": contactName, "contact_title": contactTitle, "address": address, "city": city, "region": region,
		"postal_code": postalCode, "country": country, "phone": phone})

	return result.RowsAffected, result.Error
}

///////////////////////////////////////////////////////////////////////////////////////////////
//función llamada al hacer DELETE
func removeCustomer(c echo.Context) error {

	//Se almacena el ID introducido
	CustomerID := c.QueryParam("customer_id")

	result, err := deleteCustomer(CustomerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)
	return c.String(http.StatusOK, "Cliente eliminado")
}

func deleteCustomer(customerID string) (int64, error) {

	//se conecta a la base de datos
	db := dbConnect()

	//Donde el ID sea igual, ejecuta un DELETE
	result := db.Where("customer_id = ?", customerID).Delete(&Customer{})
	return result.RowsAffected, result.Error
}
