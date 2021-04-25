package customers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/labstack/echo/v4"
)

func getCustomer(c echo.Context) error {

	CustomerID := c.QueryParam("customerID")
	log.Printf(CustomerID)
	customer, err := selectCustomer(CustomerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, customer)
}

func selectCustomer(CustomerID string) (Customer, error) {

	log.Printf("CONECTANDO...\n")
	connectionSql()
	log.Printf("TERMINADO DE CONECTAR\n")

	customer := Customer{}

	tsql := fmt.Sprintf("SELECT * FROM Customers WHERE CustomerID = '%s';", CustomerID)
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error al leer las filas " + err.Error())
		return customer, err
	}
	defer rows.Close()

	for rows.Next() {
		er := rows.Scan(&customer.CustomerID, &customer.CompanyName, &customer.ContactName, &customer.ContactTitle,
			&customer.Address, &customer.City, &customer.Region, &customer.PostalCode,
			&customer.Country, &customer.Phone, &customer.Fax)
		if err != nil {
			fmt.Println("Error al leer las filas: " + err.Error())
			return customer, er
		}

	}
	defer db.Close()
	return customer, nil
}

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
		customer.Country, customer.Phone, customer.Fax)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)

	log.Printf("Este es un cliente: %#v", customer)

	return c.String(http.StatusOK, "tenemos tu cliente")
}

func insertCustomer(customerID string, companyName string, contactName string, contactTitle string, address string, city string,
	region string, postalCode string, country string, phone string, fax string) (int64, error) {

	log.Printf("CONECTANDO...\n")
	connectionSql()
	log.Printf("TERMINADO DE CONECTAR\n")

	tsql := fmt.Sprintf("INSERT INTO Customers VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');",
		customerID, companyName, contactName, contactTitle, address, city, region, postalCode, country, phone, fax)

	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error al insertar nueva fila " + err.Error())
		return -1, err
	}

	defer db.Close()
	return result.RowsAffected()
}

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
		customer.Country, customer.Phone, customer.Fax)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)

	log.Printf("Este es un cliente: %#v", customer)

	return c.String(http.StatusOK, "tenemos tu cliente")

}

func updateCustomer(customerID string, companyName string, contactName string, contactTitle string, address string, city string,
	region string, postalCode string, country string, phone string, fax string) (int64, error) {

	log.Printf("CONECTANDO...\n")
	connectionSql()
	log.Printf("TERMINADO DE CONECTAR\n")

	tsql := fmt.Sprintf("UPDATE Customers SET CompanyName='%s',ContactName = '%s',ContactTitle='%s',Address='%s',City='%s',Region='%s',PostalCode='%s',Country='%s',Phone='%s',Fax='%s' WHERE CustomerID = '%s';",
		companyName, contactName, contactTitle, address, city, region, postalCode, country, phone, fax, customerID)

	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error al actualizar la fila " + err.Error())
		return -1, err
	}

	defer db.Close()
	return result.RowsAffected()
}

func removeCustomer(c echo.Context) error {
	CustomerID := c.QueryParam("customerID")

	result, err := deleteCustomer(CustomerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)
	return c.String(http.StatusOK, "Cliente eliminado")
}

func deleteCustomer(customerID string) (int64, error) {
	log.Printf("CONECTANDO...\n")
	connectionSql()
	log.Printf("TERMINADO DE CONECTAR\n")

	tsql := fmt.Sprintf("DELETE FROM Customers WHERE CustomerID = '%s';", customerID)
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
