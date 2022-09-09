package domain

type Product struct {
	ProductCode   string `json:"product_code"`
	Name          string `json:"name"`
	StockQuantity uint32 `json:"stock_quantity"`
}
