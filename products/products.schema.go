package products

type Product struct {
	ProductID       int     `json:"productID"`
	ProductName     string  `json:"productName"`
	SupplierID      int     `json:"supplierID"`
	CategoryID      int     `json:"categoryID"`
	QuantityPerUnit string  `json:"quantityPerUnit"`
	UnitPrice       float64 `json:"unitPrice"` //money
	UnitsInStock    int     `json:"unitsInStock"`
	UnitsOnOrder    int     `json:"unitsOnOrder"`
	ReorderLevel    int     `json:"reorderLevel"`
	Discontinued    bool    `json:"discontinued"`
}
