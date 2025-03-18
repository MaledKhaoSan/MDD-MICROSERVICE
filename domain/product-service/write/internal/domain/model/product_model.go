package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductCategory struct {
	CategoryID   uuid.UUID `json:"category_id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CategoryName string    `json:"category_name" gorm:"type:varchar(100);unique;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Product struct {
	ProductID    uuid.UUID `json:"product_id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ProductName  string    `json:"product_name" gorm:"type:varchar(255);not null"`
	ProductDesc  string    `json:"product_description" gorm:"type:text"`
	ProductPrice float64   `json:"product_price" gorm:"type:decimal(10,2);not null;check:product_price >= 0"`
	CategoryID   uuid.UUID `json:"category_id" gorm:"type:uuid;index"`
	StoreID      uuid.UUID `json:"store_id" gorm:"type:uuid;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
