package event

import (
	"encoding/json"
	"log"

	"github.com/MD-PROJECT/INVENTORY-DETAIL-SERVICE/internal/model"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// ✅ Struct สำหรับรับ Payload
type StockDecreasePayload struct {
	ProductID     string `json:"product_id" validate:"required,uuid"`
	OrderQuantity int    `json:"order_quantity" validate:"required,gte=1"`
}

// ✅ Kafka Consumer - ลด Stock เมื่อมี Order ถูกสร้าง
func StockDecreaseOrderCreatedEvent(db *gorm.DB, message []byte) error {
	var req StockDecreasePayload
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

	// ✅ ค้นหา Inventory ที่มี `product_id`
	var inventory model.Inventory
	if err := db.Where("product_id = ?", req.ProductID).First(&inventory).Error; err != nil {
		log.Printf("❌ Inventory for product %s not found: %v", req.ProductID, err)
		return err
	}

	// ✅ ตรวจสอบว่าสามารถลด Stock ได้หรือไม่
	if inventory.InventoryQuantity < req.OrderQuantity {
		log.Printf("❌ Not enough stock for product %s: Available %d, Requested %d", req.ProductID, inventory.InventoryQuantity, req.OrderQuantity)
		return nil // ❌ ไม่สามารถลด Stock ได้ ไม่ต้องคืน Error เพราะอาจเป็นธุรกิจปกติ
	}

	// ✅ ลดจำนวน Stock
	inventory.InventoryQuantity -= req.OrderQuantity

	// ✅ บันทึกการเปลี่ยนแปลง
	if err := db.Save(&inventory).Error; err != nil {
		log.Printf("❌ Failed to decrease inventory quantity for product %s: %v", req.ProductID, err)
		return err
	}

	log.Printf("✅ Stock decreased for product %s by %d, new stock: %d", req.ProductID, req.OrderQuantity, inventory.InventoryQuantity)
	return nil
}
