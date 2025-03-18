package command

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MD-PROJECT/INVENTORY-EVENTS-SOURCING/internal/infra"
	"github.com/MD-PROJECT/INVENTORY-EVENTS-SOURCING/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InBoundInventoryRequest struct {
	QuantityChange int `json:"quantity_change" validate:"required"`
}

type InBoundInventoryPayload struct {
	InventoryID    string    `json:"inventory_id"`
	QuantityChange int       `json:"quantity_change"`
	Timestamp      time.Time `json:"timestamp"`
}

func InBoundInventoryEndpoint(c *fiber.Ctx, db *gorm.DB) error {
	var req InBoundInventoryRequest
	validate := validator.New()
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": " Validation failed: " + err.Error()})
	}
	inventory_id := c.Params("inventory_id")

	event_payload, _ := json.Marshal(InBoundInventoryPayload{
		InventoryID:    inventory_id,
		QuantityChange: req.QuantityChange,
		Timestamp:      time.Now(),
	})

	event := model.Event_Stores{
		AggregateID:  &inventory_id,
		EventType:    "InventoryInBoundQuantityEvent",
		EventPayload: string(event_payload),
	}

	result := db.Create(&event)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create event"})
	}

	topic := event.EventType
	infra.PublishToKafka(topic, event_payload)

	log.Println("✅ InventoryInBoundQuantityEvent published to Kafka successfully!")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":                   "Inventory event stored and published successfully",
		"inventory_id":              event.EventID,
		"inventory_quantity_change": req.QuantityChange,
	})
}
