package event

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MD-PROJECT/ORDER-HISTORY-SERVICE/internal/model"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OrderCreatedStruct struct {
	OrderID       string    `json:"order_id" validate:"required"`
	CustomerID    string    `json:"customer_id" validate:"omitempty"`
	StoreID       string    `json:"store_id" validate:"omitempty"`
	ProductID     string    `json:"product_id" validate:"omitempty"`
	OrderStatus   string    `json:"order_status" validate:"omitempty"`
	OrderPrice    float64   `json:"order_price" validate:"omitempty"`
	OrderQuantity int       `json:"order_quantity" validate:"omitempty"`
	OrderDetails  string    `json:"order_details" validate:"omitempty"`
	CreatedAt     time.Time `json:"created_at" validate:"omitempty"`
	UpdatedAt     time.Time `json:"updated_at" validate:"omitempty"`
}

func OrderCreatedEvent(db *gorm.DB, message []byte) error {
	log.Printf("✅ Consumer Get it:")
	var req OrderCreatedStruct
	validate := validator.New()

	// ✅ Unmarshal JSON payload
	if err := json.Unmarshal(message, &req); err != nil {
		log.Printf("❌ Error unmarshalling event: %v", err)
		return err
	}

	// ✅ Validate Payload
	if err := validate.Struct(&req); err != nil {
		log.Printf("❌ Validation failed: %v", err)
		return err
	}

	// ✅ ใช้ `db.Save()` เพื่อรองรับทั้ง Insert และ Update
	order := model.Orders{
		OrderID:       req.OrderID,
		CustomerID:    req.CustomerID,
		StoreID:       req.StoreID,
		ProductID:     req.ProductID,
		OrderStatus:   req.OrderStatus,
		OrderPrice:    req.OrderPrice,
		OrderQuantity: req.OrderQuantity,
		OrderDetails:  req.OrderDetails,
		CreatedAt:     req.CreatedAt,
		UpdatedAt:     req.UpdatedAt,
	}

	if err := db.Save(&order).Error; err != nil {
		log.Printf("❌ Failed to save order: %v", err)
		return err
	}

	log.Printf("✅ Order Updated: %+v\n", order)
	return nil
}
