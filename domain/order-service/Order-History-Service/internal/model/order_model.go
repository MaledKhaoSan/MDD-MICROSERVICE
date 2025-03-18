// internal/domain/model/order_model.go
package model

import (
	"time"
)

type Orders struct {
	OrderID       string    `json:"order_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	StoreID       string    `json:"store_id" gorm:"type:uuid;not null"`
	ProductID     string    `json:"product_id" gorm:"type:uuid;not null"`
	CustomerID    string    `json:"customer_id" gorm:"type:uuid;not null"`
	OrderStatus   string    `json:"order_status" gorm:"type:varchar(8);not null;check:order_status IN ('delivery', 'waiting', 'paid')"`
	OrderPrice    float64   `json:"order_price" gorm:"type:numeric(10,2);not null;default:0"`
	OrderQuantity int       `json:"order_quantity" gorm:"not null;default:0"`
	OrderDetails  string    `json:"order_details" gorm:"type:text"`
	CreatedAt     time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
