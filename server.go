package main

import (
	"Tarea3/customers"
	"Tarea3/employees"
	"Tarea3/products"
	"log"
	"net/http"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func inicio(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func dbConnect() *gorm.DB { //Función para conectarse a la base de datos
	log.Printf("CONECTANDO...\n")

	dsn := "sqlserver://@localhost:1433?database=Backend"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Printf("Connected!\n")
	return db
}

type User struct { //estructura de la tabla users
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func login(c echo.Context) error {
	username := c.QueryParam("username") //obtiene los datos del POST y los almacena
	password := c.QueryParam("password")
	user := []User{}

	log.Printf(username)
	db := dbConnect()

	//busca el usuario en la base de datos
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {

		return c.JSON(http.StatusBadRequest, "Usuario no existe")
	}

	if password != user[0].Password {
		return c.JSON(http.StatusUnauthorized, "Contraseña incorrecta")
	}

	token := jwt.New(jwt.SigningMethodHS256) //se crea un nuevo token jwt con los datos del usuario

	claims := token.Claims.(jwt.MapClaims)
	log.Printf(user[0].Username)

	claims["username"] = user[0].Username
	claims["first_name"] = user[0].FirstName
	claims["last_name"] = user[0].LastName
	claims["email"] = user[0].Email
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	t, err := token.SignedString([]byte("pia")) //se hace salting con la palabra pia
	if err != nil {
		return err
	}

	/*

		c.SetCookie(&http.Cookie{
			Name:    "token",
			Value:   t,
			Expires: time.Now().Add(time.Hour * 48),
		})*/

	return c.JSON(http.StatusOK, map[string]string{ //regresa el token como cadena
		"token": t,
	})
}

func main() {

	e := echo.New()

	//Configuración de los CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/login", login)

	//Configuración de las rutas
	customers.Routes(e)
	employees.Routes(e)
	products.Routes(e)

	//inicia el servidor
	e.Logger.Fatal(e.Start(":1323"))

}
