package inventory_service

type OrderCreatedEventModel struct {
	OrderID   string `json:"orderId"`
	ProductID int    `json:"productId"`
	Quantity  int    `json:"quantity"`
}
