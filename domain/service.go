package domain

import "github.com/edihoxhalli/gotstock/db"

func AddProduct(p *Product) *Product {
	// save
	db.Connection()
	return &Product{
		ProductCode:   "111",
		Name:          "Lesh111",
		StockQuantity: 111,
	}
}
func GetProduct(pcode string) *Product {
	// get
	db.Connection()
	return &Product{
		ProductCode:   "111",
		Name:          "Lesh111",
		StockQuantity: 111,
	}
}
func GetAll() *[]Product {
	// get ALL
	db.Connection()
	return &[]Product{
		{
			ProductCode:   "111",
			Name:          "Lesh111",
			StockQuantity: 111,
		},
		{
			ProductCode:   "222",
			Name:          "Lesh222",
			StockQuantity: 222,
		},
	}
}
func UpdateProduct(p *Product) *Product {
	// update
	db.Connection()
	return &Product{
		ProductCode:   "222",
		Name:          "Lesh222",
		StockQuantity: 222,
	}
}
func DeleteProduct(pcode string) {
	// delete
	db.Connection()
}
