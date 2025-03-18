package inventory_service

import (
	"log"
)

// ลด Stock
func InventoryDecrease(eventData []byte) {
	log.Printf("📦 [Inventory] Decrease stock received! \n\nEventData: %s", eventData)
	// var event OrderCreatedEventModel
	// if err := json.Unmarshal(eventData, &event); err != nil {
	// 	log.Printf("❌ Failed to parse OrderCreatedEvent: %v", err)
	// 	return
	// }
	// payload, _ := json.Marshal(map[string]interface{}{
	// 	"orderId":   event.OrderID,
	// 	"productId": event.ProductID,
	// 	"quantity":  event.Quantity,
	// })

	// utils.RetryHTTPPost("http://localhost:8086/inventory/decreaseInventoryStock", payload, 3)
}
