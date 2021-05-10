package products

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

func getProduct(c echo.Context) error {

	ProductID := c.QueryParam("product_id")
	log.Printf(ProductID)
	products, err := selectProduct(ProductID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, products)
}

func selectProduct(ProductID string) ([]Product, error) {

	db := dbConnect()

	products := []Product{}

	if ProductID != "" {
		db.Where("product_id = ?", ProductID).First(&products)

		return products, nil
	} else {

		result := db.Order("product_name").Find(&products)
		return products, result.Error
	}

}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
func addProduct(c echo.Context) error {
	product := Product{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&product)
	if err != nil {
		log.Printf("fallo al procesar addProduct: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	result, err := insertProduct(product.ProductID, product.ProductName, product.QuantityPerUnit, product.UnitPrice,
		product.UnitsInStock, product.Discontinued)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)

	log.Printf("Este es un empleado: %#v", product)

	return c.String(http.StatusOK, "tenemos tu producto")
}

func insertProduct(productID int, productName string, quantityPerUnit string, unitPrice float64,
	unitsInStock int, discontinued bool) (int64, error) {

	db := dbConnect()
	/*
		var disc int
		if discontinued {
			disc = 1
		} else {
			disc = 0
		}
	*/
	product := Product{ProductID: productID, ProductName: productName, QuantityPerUnit: quantityPerUnit,
		UnitPrice: unitPrice, UnitsInStock: unitsInStock, Discontinued: discontinued}
	result := db.Create(&product)

	return result.RowsAffected, result.Error

}

/////////////////////////////////////////////////////////////////////////

func putProduct(c echo.Context) error {
	product := Product{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&product)
	if err != nil {
		log.Printf("fallo al procesar putProduct: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	result, err := updateProduct(product.ProductID, product.ProductName, product.QuantityPerUnit, product.UnitPrice,
		product.UnitsInStock, product.Discontinued)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)

	log.Printf("Este es un producto: %#v", product)

	return c.String(http.StatusOK, "tenemos tu producto")

}

func updateProduct(productID int, productName string, quantityPerUnit string, unitPrice float64,
	unitsInStock int, discontinued bool) (int64, error) {

	db := dbConnect()

	result := db.Model(&Product{}).Where("product_id = ?", productID).Updates(map[string]interface{}{
		"product_name": productName, "quantity_per_unit": quantityPerUnit, "unit_price": unitPrice,
		"units_in_stock": unitsInStock, "discontinued": discontinued})

	return result.RowsAffected, result.Error
}

///////////////////////////////////////////////////////////////////////////////////////////////
func removeProduct(c echo.Context) error {
	ProductID := c.QueryParam("product_id")

	result, err := deleteProduct(ProductID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)
	return c.String(http.StatusOK, "Producto eliminado")
}

func deleteProduct(productID string) (int64, error) {

	db := dbConnect()

	result := db.Where("product_id = ?", productID).Delete(&Product{})
	return result.RowsAffected, result.Error
}
