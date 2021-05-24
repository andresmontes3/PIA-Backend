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

//función que se llama al hacer GET
func getProduct(c echo.Context) error {

	ProductID := c.QueryParam("product_id")
	log.Printf(ProductID)
	products, err := selectProduct(ProductID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, products) //regresa un json con los productos
}

func selectProduct(ProductID string) ([]Product, error) {
	//se conecta a la base de datos
	db := dbConnect()

	products := []Product{}

	//Si se ingresó un ID, se busca ese ID, de lo contrario se muestra la lista completa
	if ProductID != "" {
		db.Where("product_id = ?", ProductID).First(&products)

		return products, nil
	} else {

		result := db.Order("product_name").Find(&products)
		return products, result.Error
	}

}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

//función que se llama al hacer POST
func addProduct(c echo.Context) error {
	product := Product{}

	defer c.Request().Body.Close()
	//obtiene los datos del JSON y los almacena en employee de tipo Product

	err := json.NewDecoder(c.Request().Body).Decode(&product)
	if err != nil {
		log.Printf("fallo al procesar addProduct: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	//Se envian a una función para trabajar con ellos en la base de datos
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
	//se conecta a la base de datos
	db := dbConnect()

	//se crea un nuevo product con los datos recibidos y se insertan con el db.Create
	product := Product{ProductID: productID, ProductName: productName, QuantityPerUnit: quantityPerUnit,
		UnitPrice: unitPrice, UnitsInStock: unitsInStock, Discontinued: discontinued}
	result := db.Create(&product)

	return result.RowsAffected, result.Error

}

/////////////////////////////////////////////////////////////////////////
//función llamada al hacer PUT
func putProduct(c echo.Context) error {
	product := Product{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&product)
	if err != nil {
		log.Printf("fallo al procesar putProduct: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	//Al igual que en post se reciben todos los parámetros a través de un JSON
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
	//se conecta a la base de datos
	db := dbConnect()

	//se hace el UPDATE con un WHERE donde el ID introducido sea el mismo
	result := db.Model(&Product{}).Where("product_id = ?", productID).Updates(map[string]interface{}{
		"product_name": productName, "quantity_per_unit": quantityPerUnit, "unit_price": unitPrice,
		"units_in_stock": unitsInStock, "discontinued": discontinued})

	return result.RowsAffected, result.Error
}

///////////////////////////////////////////////////////////////////////////////////////////////
//función llamada al hacer DELETE
func removeProduct(c echo.Context) error {

	//Se almacena el ID introducido
	ProductID := c.QueryParam("product_id")

	result, err := deleteProduct(ProductID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Printf("Filas modificadas: %d\n", result)
	return c.String(http.StatusOK, "Producto eliminado")
}

func deleteProduct(productID string) (int64, error) {
	//se conecta a la base de datos
	db := dbConnect()

	//Donde el ID sea igual, ejecuta un DELETE
	result := db.Where("product_id = ?", productID).Delete(&Product{})
	return result.RowsAffected, result.Error
}
