package products

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/labstack/echo/v4"
)

func getProduct(c echo.Context) error {

	ProductID := c.QueryParam("productID")

	Product, err := selectProduct(ProductID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, Product)
}

func selectProduct(ProductID string) (Product, error) {

	log.Printf("CONECTANDO...\n")
	connectionSql()
	log.Printf("TERMINADO DE CONECTAR\n")

	Product := Product{}

	tsql := fmt.Sprintf("SELECT * FROM Products WHERE ProductID = %s;", ProductID)
	log.Printf(ProductID)
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error al leer las filas " + err.Error())
		return Product, err
	}
	defer rows.Close()

	for rows.Next() {
		er := rows.Scan(&Product.ProductID, &Product.ProductName, &Product.SupplierID, &Product.CategoryID, &Product.QuantityPerUnit,
			Product.UnitPrice, &Product.UnitsInStock, &Product.UnitsOnOrder, &Product.ReorderLevel, &Product.Discontinued)
		if err != nil {
			fmt.Println("Error al leer las filas: " + err.Error())
			return Product, er
		}

	}
	defer db.Close()
	return Product, nil
}

func putProduct(c echo.Context) error {
	Product := Product{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&Product)
	if err != nil {
		log.Printf("fallo al procesar addProduct: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	result, err := updateProduct(Product.ProductID, Product.ProductName, Product.SupplierID, Product.CategoryID, Product.QuantityPerUnit,
		Product.UnitPrice, Product.UnitsInStock, Product.UnitsOnOrder, Product.ReorderLevel, Product.Discontinued)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)

	log.Printf("Este es un cliente: %#v", Product)

	return c.String(http.StatusOK, "tenemos tu cliente")

}

func updateProduct(ProductID int, ProductName string, SupplierID int, CategoryID int, QuantityPerUnit string,
	UnitPrice float64, UnitsInStock int, UnitsOnOrder int, ReorderLevel int, Discontinued bool) (int64, error) {

	log.Printf("CONECTANDO...\n")
	connectionSql()
	log.Printf("TERMINADO DE CONECTAR\n")
	var discontinued int
	if Discontinued {
		discontinued = 1
	} else {
		discontinued = 0
	}
	tsql := fmt.Sprintf("UPDATE Products SET ProductName='%s',SupplierID=%d,CategoryID=%d,QuantityPerUnit='%s',UnitPrice=%f,UnitsInStock=%d,UnitsOnOrder=%d,ReorderLevel=%d,Discontinued='%d' WHERE ProductID = %d;",
		ProductName, SupplierID, CategoryID, QuantityPerUnit,
		UnitPrice, UnitsInStock, UnitsOnOrder, ReorderLevel, discontinued, ProductID)

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
