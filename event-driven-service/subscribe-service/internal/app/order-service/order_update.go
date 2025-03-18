package order_service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MD-PROJECT/SUBSCRIBE-SERVICE/internal/app/utils"
)

type OrderUpdateEventModel struct {
	OrderID       string    `json:"order_id" validate:"required"`
	CustomerID    string    `json:"customer_id" validate:"required"`
	StoreID       string    `json:"store_id" validate:"required"`
	OrderStatus   string    `json:"order_status" validate:"required"`
	OrderPrice    float64   `json:"order_price" validate:"gte=0,required"`
	OrderQuantity int       `json:"order_quantity" validate:"gt=0,required"`
	OrderDetails  string    `json:"order_details" validate:"omitempty"`
	CreatedAt     time.Time `json:"created_at" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at" validate:"required"`
}

func OrderUpdate(eventData []byte) {
	log.Printf("📦 [Order] OrderUpdate! \n\nEventData: %s", eventData)

	var event OrderUpdateEventModel
	if err := json.Unmarshal(eventData, &event); err != nil {
		log.Printf("❌ Failed to parse OrderUpdateEvent: %v", err)
		return
	}

	payload, _ := json.Marshal(event)
	utils.RetryHTTP("PUT", "http://localhost:8181/orders/api", payload, 3)

}
