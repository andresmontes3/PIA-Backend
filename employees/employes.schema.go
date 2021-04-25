package employees

type Employee struct {
	EmployeeID      int    `json:"employeeID"`
	LastName        string `json:"lastName"`
	FirstName       string `json:"firstName"`
	Title           string `json:"title"`
	TitleOfCourtesy string `json:"titleOfCourtesy"`
	BirthDate       string `json:"birthDate"`
	HireDate        string `json:"hireDate"`
	Address         string `json:"address"`
	City            string `json:"city"`
	Region          string `json:"region"`
	PostalCode      string `json:"postalCode"`
	Country         string `json:"country"`
	HomePhone       string `json:"homePhone"`
	Extension       string `json:"extension"`
	Photo           string `json:"photo"` //img?
	Notes           string `json:"notes"`
	ReportsTo       int    `json:"reportsTo"`
	PhotoPath       string `json:"photoPath"`
}
