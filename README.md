
# PIA Backend
#### Integrantes
##### Andrés Isaac Montes Bartolo 1854017
----
## Base de datos
1. Se requiere tener instalado [Microsoft SQL Server](https://www.microsoft.com/es-mx/sql-server/sql-server-downloads) y el [SSMS](https://docs.microsoft.com/en-us/sql/ssms/download-sql-server-management-studio-ssms?view=sql-server-ver15) para poder ejecutar el archivo backendBdd.sql y crear la base de datos
2. Se ejecuta el archivo desde SSMS

## Instalación
1. Para instalación del backend, leer [readme](https://github.com/andresmontes3/PIA-Backend/tree/master/backend#readme)


## Requests en Postman
* Importar [PIA.postman_collection.json](https://github.com/andresmontes3/PIA-Backend/blob/master/PIA.postman_collection.json ) en Postman
* Esta colección contiene todos los endpoint de la API con información predefinida
* Primero se ejecuta el request Login para obtener el bearer token, este se reemplaza **manualmente** en el header de **Authorization** de cada request desde Postman
* Para este proyecto existe un solo usuario que puede hacer uso de la API, por lo que si se intenta ingresar con otro usuario o con una contraseña diferente, el servidor devolverá error. Este cuenta con las datos:
```
	username = josuefdz
	password = contra
```
* Peticiones:
    * **POST Login:** http://localhost:1323/login?username="usuario"&password="contraseña" </br>
    

    Regresa:
    ```
    {"token": <nuevo Token> }
    ```
   
    * **GET Customers:** http://localhost:1323/api/customers </br>
    Retorna: 
    Lista de clientes
   * **GET Customer:** http://localhost:1323/api/customers?customer_id="id" </br>
    Retorna: 
    Cliente
    *  **GET Employees:** http://localhost:1323/api/employees </br>
    Retorna: 
    Lista de empleados
    * **GET Employee:** http://localhost:1323/api/employees?employee_id="id" </br>
    Retorna:
    Empleado
     *  **GET Products:** http://localhost:1323/api/products </br>
    Retorna: 
    Lista de productos
     * **GET Product:** http://localhost:1323/api/products?product_id="id" </br>
    Retorna: 
    Producto

    * **POST Customer** http://localhost:1323/api/customers </br>
    Body:
    ```
    {
    "customer_id"="RIPAZ",
    "company_name": "Enchiladas Josue",
    "contact_name": "Josue Rodríguez",
    "contact_title": "Dueño",
    "address": " Calle Pinos #3451, Colonia Bosque Verde, Monterrey",
    "city": "Nuevo León",
    "region": "Centro",
    "postal_code": "64210",
    "country": "México",
    "phone": "8144393"
    }
    ```
   * **POST Employee** http://localhost:1323/api/employees </br>
    Body:
    ```
    {
    "employee_id": 1,
    "last_name": "Sánchez",
    "first_name": "Alejandra",
    "title": "Gerente",
    "hire_date": "1900-01-01T00:00:00Z",
    "address": "Calle Esmeralda #345, Colonia Piedras Preciosas, Guadalupe",
    "city": "Nuevo León",
    "region": " ",
    "postal_code": "64860",
    "country": "",
    "phone": "8168450"
    }
    ``` 
   * **POST Product** http://localhost:1323/api/products </br>
    Body:
    ```
    {
    "product_id": 3,
    "product_name": "Chespiritortas",
    "quantity_per_unit": "1",
    "unit_price": 80,
    "units_in_stock": 0,
    "discontinued": true
    }
    ```
    * **PUT Customer** http://localhost:1323/api/products </br>
    Body:
    ```
    {
    "customer_id"="RIPAZ",
    "company_name": "Enchiladas Josue2",
    "contact_name": "Josue Rodríguez",
    "contact_title": "Dueño",
    "address": " Calle Pinos #3451, Colonia Bosque Verde, Monterrey",
    "city": "Nuevo León",
    "region": "Centro",
    "postal_code": "64210",
    "country": "México",
    "phone": "8144393"
    }
    ```
    * **PUT Employee** http://localhost:1323/api/products </br>
    Body:
    ```
    {
    "employee_id": 1,
    "last_name": "Sánchez",
    "first_name": "Alexandra",
    "title": "Gerente",
    "hire_date": "1900-01-01T00:00:00Z",
    "address": "Calle Esmeralda #345, Colonia Piedras Preciosas, Guadalupe",
    "city": "Nuevo León",
    "region": " ",
    "postal_code": "64860",
    "country": "",
    "phone": "8168450"
    }
    ```
    * **PUT Product** http://localhost:1323/api/products </br>
    Body:
    ```
    {
    "product_id": 3,
    "product_name": "Chespiritortas",
    "quantity_per_unit": "8",
    "unit_price": 80,
    "units_in_stock": 0,
    "discontinued": true
    }
    ```
  *  **DELETE Customers:** http://localhost:1323/api/customers?customer_id="id" </br>

  *  **DELETE Employees:** http://localhost:1323/api/employees?employee_id="id" </br>

  *  **DELETE Products:** http://localhost:1323/api/products?product_id="id" </br>


