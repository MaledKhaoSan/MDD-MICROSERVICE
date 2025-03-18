package event

import (
	"encoding/json"
	"log"

	"github.com/MD-PROJECT/ORDER-HISTORY-SERVICE/internal/model"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// ✅ Struct ต้องตรงกับ JSON ที่ Write Model ส่งมา
type OrderUpdateStatusPayload struct {
	OrderID     string `json:"order_id" validate:"required"`
	OrderStatus string `json:"order_status" validate:"required"`
}

// ✅ Kafka Consumer - อัปเดตสถานะของ Order ใน Read Database
func OrderUpdateStatusEvent(db *gorm.DB, message []byte) error {
	var req OrderUpdateStatusPayload
	validate := validator.New()

	// ✅ Unmarshal JSON Payload
	if err := json.Unmarshal(message, &req); err != nil {
		log.Printf("❌ Error unmarshalling event: %v", err)
		return err
	}

	// ✅ Validate Payload
	if err := validate.Struct(&req); err != nil {
		log.Printf("❌ Validation failed: %v", err)
		return err
	}

	// ✅ ค้นหา Order ใน Read Database
	var order model.Orders
	if err := db.Where("order_id = ?", req.OrderID).First(&order).Error; err != nil {
		log.Printf("❌ Order not found: %v", err)
		return err
	}

	// ✅ อัปเดตเฉพาะ OrderStatus
	order.OrderStatus = req.OrderStatus

	// ✅ บันทึกการเปลี่ยนแปลง
	if err := db.Save(&order).Error; err != nil {
		log.Printf("❌ Failed to update order status: %v", err)
		return err
	}

	log.Printf("✅ Order %s status updated to %s", req.OrderID, req.OrderStatus)
	return nil
}
