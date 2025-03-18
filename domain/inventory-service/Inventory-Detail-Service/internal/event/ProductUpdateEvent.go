package event

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MD-PROJECT/INVENTORY-DETAIL-SERVICE/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductUpdateStruct struct {
	ProductID          uuid.UUID  `json:"product_id" validate:"required"`
	ProductName        string     `json:"product_name" validate:"required"`
	ProductDescription string     `json:"product_description"`
	ProductPrice       float64    `json:"product_price" validate:"required,gt=0"`
	CategoryID         *uuid.UUID `json:"category_id"`
	StoreID            uuid.UUID  `json:"store_id" validate:"required"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

func ProductUpdateEvent(db *gorm.DB, message []byte) error {
	log.Println("✅ Consumer received ProductUpdateEvent")

	var req ProductUpdateStruct
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

	// ✅ บันทึกลง Database
	product := model.Product{
		ProductID:          req.ProductID,
		ProductName:        req.ProductName,
		ProductDescription: req.ProductDescription,
		ProductPrice:       req.ProductPrice,
		CategoryID:         req.CategoryID,
		StoreID:            req.StoreID,
		CreatedAt:          req.CreatedAt,
		UpdatedAt:          req.UpdatedAt,
	}

	if err := db.Save(&product).Error; err != nil {
		log.Printf("❌ Failed to save product: %v", err)
		return err
	}

	log.Printf("✅ Product updated successfully: %+v\n", product)
	return nil
}
