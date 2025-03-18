package command

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MD-PROJECT/PRODUCT-EVENTS-SOURCING/internal/infra"
	"github.com/MD-PROJECT/PRODUCT-EVENTS-SOURCING/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UpdateProductRequest struct {
	ProductName        *string    `json:"product_name,omitempty"`
	ProductDescription *string    `json:"product_description,omitempty"`
	ProductPrice       *float64   `json:"product_price,omitempty"`
	CategoryID         *uuid.UUID `json:"category_id,omitempty"`
	StoreID            uuid.UUID  `json:"store_id" validate:"required"`
}

type UpdateProductPayload struct {
	ProductID          string     `json:"product_id"`
	ProductName        *string    `json:"product_name,omitempty"`
	ProductDescription *string    `json:"product_description,omitempty"`
	ProductPrice       *float64   `json:"product_price,omitempty"`
	CategoryID         *uuid.UUID `json:"category_id,omitempty"`
	StoreID            uuid.UUID  `json:"store_id" validate:"required"`
	Timestamp          time.Time  `json:"timestamp"`
}

func UpdateProductEndpoint(c *fiber.Ctx, db *gorm.DB) error {
	productID := c.Params("product_id")
	var req UpdateProductRequest
	validate := validator.New()

	// ✅ ตรวจสอบ JSON Payload
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed: " + err.Error()})
	}

	// ✅ สร้าง Event Payload เฉพาะค่าที่มี
	eventPayload, _ := json.Marshal(UpdateProductPayload{
		ProductID:          productID,
		ProductName:        req.ProductName,
		ProductDescription: req.ProductDescription,
		ProductPrice:       req.ProductPrice,
		CategoryID:         req.CategoryID,
		StoreID:            req.StoreID,
		Timestamp:          time.Now(),
	})

	// ✅ บันทึก Event ลง Event Store
	event := model.Event_Stores{
		AggregateID:  &productID,
		EventType:    "ProductUpdateEvent",
		EventPayload: string(eventPayload),
	}

	if err := db.Create(&event).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create event"})
	}

	infra.PublishToKafka(event.EventType, eventPayload)

	log.Println("✅ ProductUpdateEvent published to Kafka successfully!")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Product update event stored and published successfully",
		"product_id": event.AggregateID,
		"product":    event.EventPayload,
	})
}
