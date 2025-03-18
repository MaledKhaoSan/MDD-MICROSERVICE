package model

import (
	"time"

	"github.com/google/uuid"
)

type Event_Stores struct {
	EventID      string    `json:"event_id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	AggregateID  *string   `json:"aggregate_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	EventType    string    `json:"event_type" gorm:"type:varchar(50);not null"`
	EventPayload string    `json:"event_payload" gorm:"type:text;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// ProductCategory represents the product_category table
type ProductCategory struct {
	CategoryID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"category_id"`
	CategoryName string    `gorm:"type:varchar(100);not null;unique" json:"category_name"`
	CreatedAt    time.Time `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not null;default:current_timestamp" json:"updated_at"`
}

// Product represents the product table
type Product struct {
	ProductID          uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"product_id"`
	ProductName        string          `gorm:"type:varchar(255);not null" json:"product_name"`
	ProductDescription string          `gorm:"type:text" json:"product_description"`
	ProductPrice       float64         `gorm:"type:decimal(10,2);not null;check:product_price >= 0" json:"product_price"`
	CategoryID         *uuid.UUID      `gorm:"type:uuid;references:product_category(category_id);on delete set null" json:"category_id"`
	StoreID            uuid.UUID       `gorm:"type:uuid;not null" json:"store_id"`
	CreatedAt          time.Time       `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt          time.Time       `gorm:"not null;default:current_timestamp" json:"updated_at"`
	Category           ProductCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

// Location represents the location table
type Location struct {
	LocationID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"location_id"`
	Address    string    `gorm:"type:varchar(255);not null" json:"address"`
	City       *string   `gorm:"type:varchar(100)" json:"city"`
	State      *string   `gorm:"type:varchar(100)" json:"state"`
	Country    *string   `gorm:"type:varchar(100)" json:"country"`
	Latitude   *float64  `gorm:"type:double precision" json:"latitude"`
	Longitude  *float64  `gorm:"type:double precision" json:"longitude"`
	CreatedAt  time.Time `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"not null;default:current_timestamp" json:"updated_at"`
}

// Store represents the store table
type Store struct {
	StoreID    uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"store_id"`
	StoreName  string     `gorm:"type:varchar(255);not null" json:"store_name"`
	LocationID *uuid.UUID `gorm:"type:uuid;references:location(location_id);on delete set null" json:"location_id"`
	CreatedAt  time.Time  `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"not null;default:current_timestamp" json:"updated_at"`
	Location   Location   `gorm:"foreignKey:LocationID" json:"location,omitempty"`
}

// Warehouse represents the warehouse table
type Warehouse struct {
	WarehouseID   uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"warehouse_id"`
	WarehouseName string     `gorm:"type:varchar(255);not null" json:"warehouse_name"`
	LocationID    *uuid.UUID `gorm:"type:uuid;references:location(location_id);on delete set null" json:"location_id"`
	CreatedAt     time.Time  `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"not null;default:current_timestamp" json:"updated_at"`
	Location      Location   `gorm:"foreignKey:LocationID" json:"location,omitempty"`
}

// Inventory represents the inventory table
type Inventory struct {
	InventoryID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"inventory_id"`
	ProductID            uuid.UUID  `gorm:"type:uuid;not null" json:"product_id"`
	StoreID              uuid.UUID  `gorm:"type:uuid;not null;references:store(store_id);on delete cascade" json:"store_id"`
	WarehouseID          *uuid.UUID `gorm:"type:uuid;references:warehouse(warehouse_id);on delete set null" json:"warehouse_id"`
	InventoryQuantity    int        `gorm:"not null;default:0;check:inventory_quantity >= 0" json:"inventory_quantity"`
	InventoryMinQuantity int        `gorm:"column:inventory_min_qunatity;not null;default:0;check:inventory_min_qunatity >= 0" json:"inventory_min_quantity"`
	CreatedAt            time.Time  `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt            time.Time  `gorm:"not null;default:current_timestamp" json:"updated_at"`
	Product              Product    `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Store                Store      `gorm:"foreignKey:StoreID" json:"store,omitempty"`
	Warehouse            Warehouse  `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
}

// TableName overrides the table name used by Inventory to match the SQL definition
func (Inventory) TableName() string {
	return "inventory"
}

func (Event_Stores) TableName() string {
	return "inventory_event_stores"
}
