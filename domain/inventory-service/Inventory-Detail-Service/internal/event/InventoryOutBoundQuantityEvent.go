package event

import (
	"encoding/json"
	"log"

	"github.com/MD-PROJECT/INVENTORY-DETAIL-SERVICE/internal/model"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OutBoundInventoryPayload struct {
	InventoryID    string `json:"inventory_id" validate:"required,uuid"`
	QuantityChange int    `json:"quantity_change" validate:"required,gte=1"`
}

// ✅ Kafka Consumer - อัปเดต Inventory Quantity
func InventoryOutBoundQuantityEvent(db *gorm.DB, message []byte) error {
	var req OutBoundInventoryPayload
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

	// ✅ ค้นหา Inventory
	var inventory model.Inventory
	if err := db.Where("inventory_id = ?", req.InventoryID).First(&inventory).Error; err != nil {
		log.Printf("❌ Inventory not found: %v", err)
		return err
	}

	// ✅ อัปเดต Inventory Quantity
	inventory.InventoryQuantity -= req.QuantityChange

	// ✅ บันทึกการเปลี่ยนแปลง
	if err := db.Save(&inventory).Error; err != nil {
		log.Printf("❌ Failed to update inventory quantity: %v", err)
		return err
	}

	log.Printf("✅ Inventory %s quantity increased by %d, new total: %d", req.InventoryID, req.QuantityChange, inventory.InventoryQuantity)
	return nil
}
