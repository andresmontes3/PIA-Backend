
# Backend
Manual técnico

## Instalación
Primero es necesario instalar Go, para eso se descarga el msi desde este [link](https://golang.org/doc/install)
* Nos posicionamos en la carpeta backend
`cd <directorio backend>`
* Encendemos el servidor ejecutando el siguiente comando
`go run server.go`
----------

### EndPoints
* **Login**
	* **Pruebas**:
	* **URL**: http://localhost:1323/login?username=usuario&password=contraseña
	* **Metodo HTTP**: POST
	* **Parámetros del URL**: ``username`` ``password``
	* **Parámetros del body**: No aplica
El Token que regresa se cambia **manualmente** en el header de **Authorization** de los demas request como Bearer
* **Customers**
	* **GET Customers**
		* **Pruebas**:
		* **URL**: http://localhost:1323/api/customers
		* **Metodo HTTP**: GET
		* **Parámetros del URL**: No aplica
		* **Parámetros del body**: No aplica
	* **GET Customer**
		* **Pruebas**:
		* **URL**: http://localhost:1323/api/customers?customer_id=id
		* **Metodo HTTP**: GET
		* **Parámetros del URL**: ``customer_id``
		* **Parámetros del body**: No aplica
	* **POST Customer**
		 * **Pruebas**:
		*  **URL**: http://localhost:1323/api/customers
		* **Metodo HTTP**: POST
		* **Parámetros del URL**: No aplica
		* **Parámetros del body**: 
    ```
    {
    "customer_id"="id",
    "company_name": "compañia",
    "contact_name": "nombre",
    "contact_title": "cargo",
    "address": "direccion",
    "city": "ciudad",
    "region": "region",
    "postal_code": "codigo postal",
    "country": "pais",
    "phone": "numero telefonico"
    }
    ```	
    
	* **PUT Customer**
		* **Pruebas**:
		*  **URL**: http://localhost:1323/api/customers
		* **Metodo HTTP**: PUT
		* **Parámetros del URL**: No aplica
		* **Parámetros del body**: 
	 ```
    {
    "customer_id"="id",
    "company_name": "compañia",
    "contact_name": "nombre",
    "contact_title": "cargo",
    "address": "direccion",
    "city": "ciudad",
    "region": "region",
    "postal_code": "codigo postal",
    "country": "pais",
    "phone": "numero telefonico"
    }
    ```	

	* **DELETE Customer**
		* **Pruebas**:
		* **URL**: http://localhost:1323/api/customers?customer_id=id
		* **Metodo HTTP**: DELETE
		* **Parámetros del URL**: ``customer_id``
		* **Parámetros del body**: No aplica

* **Employees**
	* **GET Employees**
	* **GET Employee**
	* **POST Employee**
	* **PUT Employee**
	* **DELETE Employee**
* **Products**
	* **GET Products**
	* **GET Product**
	* **POST Product**
	* **PUT Product**
	* **DELETE Product**


 

