package command

import (
	"log"
	"time"

	"github.com/MD-PROJECT/PRODUCT-SERVICE-READ-MODEL/internal/domain/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProductUpdateEventModel represents the payload for updating a product
type ProductUpdateEventModel struct {
	ProductID    string    `json:"product_id" validate:"required"`
	ProductName  string    `json:"product_name" validate:"required"`
	ProductDesc  string    `json:"product_description" validate:"omitempty"`
	ProductPrice float64   `json:"product_price" validate:"gte=0,required"`
	CategoryID   string    `json:"category_id" validate:"required"`
	StoreID      string    `json:"store_id" validate:"required"`
	CreatedAt    time.Time `json:"created_at" validate:"required"`
	UpdatedAt    time.Time `json:"updated_at" validate:"required"`
}

// UpdateProductEndpoint updates the product details
func UpdateProductEndpoint(c *fiber.Ctx, db *gorm.DB) error {
	var req ProductUpdateEventModel
	validate := validator.New()

	// ✅ ตรวจสอบ Payload
	if err := c.BodyParser(&req); err != nil || validate.Struct(&req) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	// ✅ ใช้ `db.Save()` เพื่อรองรับทั้ง Insert และ Update
	product := model.Product{
		ProductID:    uuid.Must(uuid.Parse(req.ProductID)),
		ProductName:  req.ProductName,
		ProductDesc:  req.ProductDesc,
		ProductPrice: req.ProductPrice,
		CategoryID:   uuid.Must(uuid.Parse(req.CategoryID)),
		StoreID:      uuid.Must(uuid.Parse(req.StoreID)),
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
	}

	// Save the product, upsert logic is handled by db.Save
	if err := db.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save product"})
	}

	log.Printf("✅ Product Upserted: %+v\n", product)
	return c.Status(fiber.StatusOK).JSON(product)
}
