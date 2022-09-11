package db

type Product struct {
	ID            *int64 `json:"id,omitempty"`
	ProductCode   string `json:"product_code"`
	Name          string `json:"name"`
	StockQuantity uint32 `json:"stock_quantity"`
}
