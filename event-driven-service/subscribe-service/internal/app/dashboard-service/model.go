package dashboard_service

type DashboardOrderUpdateEventModel struct {
	OrderID     string `json:"order_id"`
	OrderStatus string `json:"order_status"`
}
