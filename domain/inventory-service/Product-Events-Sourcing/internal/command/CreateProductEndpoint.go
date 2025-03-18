package command

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/PRODUCT-EVENTS-SOURCING/internal/infra"
	"github.com/MD-PROJECT/PRODUCT-EVENTS-SOURCING/internal/model"
)

type CreateProductRequest struct {
	ProductName        string     `json:"product_name" validate:"required"`
	ProductDescription string     `json:"product_description"`
	ProductPrice       float64    `json:"product_price" validate:"required,gt=0"`
	CategoryID         *uuid.UUID `json:"category_id"`
	StoreID            uuid.UUID  `json:"store_id" validate:"required"`
}

func generateUniqueUUID(db *gorm.DB) string {
	for {
		aggregateID := uuid.NewString()
		var count int64
		db.Model(&model.Event_Stores{}).Where("aggregate_id = ?", aggregateID).Count(&count)
		if count == 0 {
			return aggregateID
		}
	}
}

func CreateProductEndpoint(c *fiber.Ctx, db *gorm.DB) error {
	var req CreateProductRequest
	validate := validator.New()

	// ✅ ตรวจสอบ JSON Payload
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed: " + err.Error()})
	}

	// ✅ Generate Unique Product ID
	productID := generateUniqueUUID(db)

	// ✅ สร้าง Event Payload
	eventPayload, _ := json.Marshal(model.Product{
		ProductID:          uuid.MustParse(productID),
		ProductName:        req.ProductName,
		ProductDescription: req.ProductDescription,
		ProductPrice:       req.ProductPrice,
		CategoryID:         req.CategoryID,
		StoreID:            req.StoreID,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	})

	// ✅ บันทึกลง Event Store
	event := model.Event_Stores{
		AggregateID:  &productID,
		EventType:    "ProductCreatedEvent",
		EventPayload: string(eventPayload),
	}

	if err := db.Create(&event).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create event"})
	}

	// ✅ Publish Event ไป Kafka
	infra.PublishToKafka(event.EventType, eventPayload)

	log.Println("✅ ProductCreatedEvent published to Kafka successfully!")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Product event stored and published successfully",
		"product_id": event.AggregateID,
		"product":    event.EventPayload,
	})
}
