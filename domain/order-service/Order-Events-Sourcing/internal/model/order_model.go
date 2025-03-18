package model

import "time"

type Event_Stores struct {
	EventID      string    `json:"event_id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	AggregateID  *string   `json:"aggregate_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	EventType    string    `json:"event_type" gorm:"type:varchar(50);not null"`
	EventPayload string    `json:"event_payload" gorm:"type:text;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type Orders struct {
	OrderID       string    `json:"order_id" validate:"required"`
	CustomerID    string    `json:"customer_id" validate:"required"`
	StoreID       string    `json:"store_id" validate:"required"`
	ProductID     string    `json:"product_id" validate:"required"`
	OrderStatus   string    `json:"order_status" validate:"required"`
	OrderPrice    float64   `json:"order_price" validate:"gte=0,required"`
	OrderQuantity int       `json:"order_quantity" validate:"gt=0,required"`
	OrderDetails  string    `json:"order_details" validate:"omitempty"`
	CreatedAt     time.Time `json:"created_at" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at" validate:"required"`
}
