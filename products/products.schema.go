package products

type Product struct {
	ProductID       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	QuantityPerUnit string  `json:"quantity_per_unit"`
	UnitPrice       float64 `json:"unit_price"` //money
	UnitsInStock    int     `json:"units_in_stock"`
	Discontinued    bool    `json:"discontinued"`
}
