package employees

type Employee struct {
	EmployeeID int    `json:"employee_id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	Title      string `json:"title"`
	HireDate   string `json:"hire_date"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Region     string `json:"region"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
}
