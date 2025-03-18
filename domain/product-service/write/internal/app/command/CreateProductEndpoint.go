package command

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/PRODUCT-SERVICE-WRITE-MODEL/internal/app/producer"
	"github.com/MD-PROJECT/PRODUCT-SERVICE-WRITE-MODEL/internal/domain/model"
)

type CreateProductRequest struct {
	ProductID    string    `json:"product_id" validate:"required"`
	ProductName  string    `json:"product_name" validate:"required"`
	ProductDesc  string    `json:"product_description" validate:"omitempty"`
	ProductPrice float64   `json:"product_price" validate:"gte=0,required"`
	CategoryID   string    `json:"category_id" validate:"required"`
	StoreID      string    `json:"store_id" validate:"required"`
	CreatedAt    time.Time `json:"created_at" validate:"required"`
	UpdatedAt    time.Time `json:"updated_at" validate:"required"`
}

func CreateProductEndpoint(c *fiber.Ctx, db *gorm.DB) error {
	var req CreateProductRequest
	validate := validator.New()

	if err := c.BodyParser(&req); err != nil || validate.Struct(&req) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

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

	if err := db.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
	}
	log.Printf("📦 Product Created: %+v\n", product)

	producer.PublishProductCreatedEvent(product)
	return c.Status(fiber.StatusCreated).JSON(product)
}
