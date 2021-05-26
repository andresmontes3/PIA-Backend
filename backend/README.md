

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
	* **Pruebas**:<br>![login](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/login.jpg)
	* **URL**: http://localhost:1323/login?username=usuario&password=contraseña
	* **Metodo HTTP**: POST
	* **Parámetros del URL**: ``username`` ``password``
	* **Parámetros del body**: No aplica<br><br>
El Token que regresa se cambia **manualmente** en el header de **Authorization** de los demas request como Bearer.
Solo existe el usuario josuefdz con contraseña=contra en la base de datos, si se intenta entrar con otro la petición será rechazada.
* **Customers**
	* **GET Customers**
		* **Pruebas**:<br>![getCustomers](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Customer/GetCustomers.png)
		* **URL**: http://localhost:1323/api/customers
		* **Metodo HTTP**: GET
		* **Parámetros del URL**: No aplica
		* **Parámetros del body**: No aplica
	* **GET Customer**
		* **Pruebas**:<br>![getCustomer](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Customer/GetCustomer.png)
		* **URL**: http://localhost:1323/api/customers?customer_id=id
		* **Metodo HTTP**: GET
		* **Parámetros del URL**: ``customer_id``
		* **Parámetros del body**: No aplica
	* **POST Customer**
		* **Pruebas**:<br>![postCustomer](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Customer/PostCustomer.png)
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
		* **Pruebas**:<br>![putCustomer](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Customer/PutCustomer.png)
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
		* **Pruebas**:<br>![deleteCustomer](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Customer/DeleteCustomer.png)<br>![deletedCustomer](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Customer/GetCustomers-AfterDelete.png)
		* **URL**: http://localhost:1323/api/customers?customer_id=id
		* **Metodo HTTP**: DELETE
		* **Parámetros del URL**: ``customer_id``
		* **Parámetros del body**: No aplica

* **Employees**
	* **GET Employees**
	    * **Pruebas**:<br>![getEmployees](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Employee/GetEmployees.png)
	    * **URL**: http://localhost:1323/api/employees
	    * **Metodo HTTP**: GET
	    * **Parámetros del URL**: No aplica
	    * **Parámetros del body**: No aplica
	* **GET Employee**
	    * **Pruebas**:<br>![getEmployee](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Employee/GetEmployee.png)
	    * **URL**: http://localhost:1323/api/employees?employee_id=id
	    * **Metodo HTTP**: GET
	    * **Parámetros del URL**: ``employee_id``
	    * **Parámetros del body**: No aplica
	* **POST Employee**
	     * **Pruebas**:<br>![postEmployee](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Employee/PostEmployee.png)
	     * **URL**: http://localhost:1323/api/employees
	     * **Metodo HTTP**: POST
	     * **Parámetros del URL**: No aplica
	     * **Parámetros del body**:
	     ```
	     {
	     "employee_id":id entero,
	     "last_name": "apellido", 
             "first_name": "nombre", 
             "title": "titulo", 
             "hire_date": "fecha contratacion",
             "address": "direccion", 
             "city": "ciudad", 
             "region": "region", 
             "postal_code": "codigo postal", 
	     "country":"pais",
             "phone": "telefono"
             }
	     ```
	* **PUT Employee**
	    * **Pruebas**:<br>![putEmployee](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Employee/PutEmployee.png)
	    * **URL**: http://localhost:1323/api/employees
	    * **Metodo HTTP**: PUT
	    * **Parámetros del URL**: No aplica
	    * **Parámetros del body**:
	    ```
            {
	    "employee_id": id entero,
	    "last_name": "apellido", 
            "first_name": "nombre", 
            "title": "titulo", 
            "hire_date": "fecha contratacion",
            "address": "direccion", 
            "city": "ciudad", 
            "region": "region", 
            "postal_code": "codigo postal", 
	    "country":"pais",
            "phone": "telefono"
            }
         ```
	* **DELETE Employee**
	    * **Pruebas**:<br>![deleteEmployee](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Employee/DeleteEmployee.png)<br>![deletedEmployee](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Employee/GetEmployee-afterDelete.png)
	    * **URL**: http://localhost:1323/api/employees?employee_id=id
	    * **Metodo HTTP**: DELETE
	    * **Parámetros del URL**: ``employee_id``
	    * **Parámetros del body**: No aplica
	
* **Products**
	* **GET Products**
	    * **Pruebas**:<br>![getProducts](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Products/getProducts.png)
	    * **URL**: http://localhost:1323/api/products
	    * **Metodo HTTP**: GET
	    * **Parámetros del URL**: No aplica
	    * **Parámetros del body**: No aplica
	* **GET Product**
	    * **Pruebas**:<br>![getProduct](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Products/getProduct.png)
	    * **URL**: http://localhost:1323/api/products?product_id=id
	    * **Metodo HTTP**: GET
	    * **Parámetros del URL**: ``product_id``
	    * **Parámetros del body**: No aplica
	* **POST Product**
	     * **Pruebas**:<br>![postProduct](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Products/postproduct.png)
	     * **URL**: http://localhost:1323/api/products
	     * **Metodo HTTP**: POST
	     * **Parámetros del URL**: No aplica
	     * **Parámetros del body**:
	```
	{
	"product_id": id entero,
	"product_name": "nombre",
	"quantity_per_unit": "cantidad por unidad",
	"unit_price": precio por unidad,
	"units_in_stock": unidades en stock,
	"discontinued": true/false
	}    
	```
	* **PUT Product**
	    * **Pruebas**:<br>![putProduct](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Products/putProduct.png)
	    * **URL**: http://localhost:1323/api/products
	    * **Metodo HTTP**: PUT
	    * **Parámetros del URL**: No aplica
	    * **Parámetros del body**:
	```
	{
	"product_id": id entero,
	"product_name": "nombre",
	"quantity_per_unit": "cantidad por unidad",
	"unit_price": precio por unidad,
	"units_in_stock": unidades en stock,
	"discontinued": true/false
	}    
	```
	* **DELETE Product**
	    * **Pruebas**:<br>![deleteProduct](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Products/deleteProduct.png)<br>![deletedProduct](https://github.com/andresmontes3/PIA-Backend/blob/master/capturas/Products/getProduct-afterdelete.png)
	    * **URL**: http://localhost:1323/api/products?product_id=id
	    * **Metodo HTTP**: DELETE
	    * **Parámetros del URL**: ``product_id``
	    * **Parámetros del body**: No aplica


 

