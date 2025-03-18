package dashboard_service

import (
	"encoding/json"
	"log"

	"github.com/MD-PROJECT/SUBSCRIBE-SERVICE/internal/app/utils"
)

// DashBoardOrderUpdateStatus - ส่ง HTTP Request ไปอัปเดตสถานะใน Dashboard Service
func DashBoardOrderUpdateStatus(eventData []byte) {
	log.Printf("📦 [Dashboard] Update DashboardOrderStatus! \n\nEventData: %s", eventData)

	var event DashboardOrderUpdateEventModel
	if err := json.Unmarshal(eventData, &event); err != nil {
		log.Printf("❌ Failed to parse OrderCreatedEvent: %v", err)
		return
	}

	payload, _ := json.Marshal(event)
	utils.RetryHTTP("PUT", "http://localhost", payload, 3)

	log.Printf("✅ Order status updated in Dashboard successfully!")
}
