package notification_service

import (
	"log"
)

type OrderEvent struct {
	OrderID string `json:"orderId"`
	ShopID  string `json:"shopId"`
}

// แจ้งเตือนคำสั่งซื้อใหม่
func NotificationNewOrder(eventData []byte) {
	log.Printf("🔔 [Notification] New order received! OrderID")
	// var event OrderEvent
	// if err := json.Unmarshal(eventData, &event); err != nil {
	// 	log.Printf("❌ Failed to parse OrderCreatedEvent: %v", err)
	// 	return
	// }

	// log.Printf("🔔 [Notification] New order received! OrderID: %s, ShopID: %s",
	// 	event.OrderID, event.ShopID)

	// TODO: ส่ง Notification ไปยังระบบที่เกี่ยวข้อง
}
