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

func getCustomer(c echo.Context) error {

	CustomerID := c.QueryParam("customer_id")
	log.Printf(CustomerID)
	customers, err := selectCustomer(CustomerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, customers)
}

func selectCustomer(CustomerID string) ([]Customer, error) {

	db := dbConnect()

	customers := []Customer{}

	if CustomerID != "" {
		db.Where("customer_id = ?", CustomerID).First(&customers)

		return customers, nil
	} else {

		result := db.Order("company_name").Find(&customers)

		return customers, result.Error
	}

}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
func addCustomer(c echo.Context) error {
	customer := Customer{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&customer)
	if err != nil {
		log.Printf("fallo al procesar addCustomer: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

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

	db := dbConnect()

	customer := Customer{CustomerID: customerID, CompanyName: companyName, ContactName: contactName,
		ContactTitle: contactTitle, Address: address, City: city, Region: region, PostalCode: postalCode,
		Country: country, Phone: phone}
	result := db.Create(&customer)

	return result.RowsAffected, result.Error

}

/////////////////////////////////////////////////////////////////////////

func putCustomer(c echo.Context) error {
	customer := Customer{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&customer)
	if err != nil {
		log.Printf("fallo al procesar addCustomer: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

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

	db := dbConnect()

	result := db.Model(&Customer{}).Where("customer_id = ?", customerID).Updates(map[string]interface{}{"company_name": companyName,
		"contact_name": contactName, "contact_title": contactTitle, "address": address, "city": city, "region": region,
		"postal_code": postalCode, "country": country, "phone": phone})

	return result.RowsAffected, result.Error
}

///////////////////////////////////////////////////////////////////////////////////////////////
func removeCustomer(c echo.Context) error {
	CustomerID := c.QueryParam("customer_id")

	result, err := deleteCustomer(CustomerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)
	return c.String(http.StatusOK, "Cliente eliminado")
}

func deleteCustomer(customerID string) (int64, error) {
	log.Printf("CONECTANDO...\n")

	db := dbConnect()

	result := db.Where("customer_id = ?", customerID).Delete(&Customer{})
	return result.RowsAffected, result.Error
}
