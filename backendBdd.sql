CREATE DATABASE Backend

use Backend

CREATE TABLE employees(
	employee_id int primary key,
	last_name nvarchar(20) not null,
	first_name nvarchar(10) not null,
	title nvarchar(30),
	hire_date datetime,
	address nvarchar(60),
	city nvarchar(15),
	region nvarchar(15),
	postal_code nvarchar(10),
	country nvarchar(15),
	phone nvarchar(24),
);

CREATE TABLE customers(
	customer_id nchar(5) primary key,
	company_name nvarchar(40) not null,
	contact_name nvarchar(30),
	contact_title nvarchar(60),
	address nvarchar(60),
	city nvarchar(15),
	region nvarchar(15),
	postal_code nvarchar(10),
	country nvarchar(15),
	phone nvarchar(24),
);

CREATE TABLE products(
	product_id int primary key,
	product_name nvarchar(40) not null,
	quantity_per_unit nvarchar(20),
	unit_price money,
	units_in_stock smallint,
	discontinued bit
);



INSERT INTO Customers VALUES ('MYID','Panadería Roberto','Roberto Fuentes','Dueño',' Calle Nogal #3556, Colonia Bosque Verde, Monterrey','Nuevo León','Centro','64210','México','8123451')


select * from customers

select * from employees

select * from products