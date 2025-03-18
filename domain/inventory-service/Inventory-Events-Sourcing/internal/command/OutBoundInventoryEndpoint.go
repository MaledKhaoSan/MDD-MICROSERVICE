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

type OutBoundInventoryRequest struct {
	QuantityChange int `json:"quantity_change" validate:"required"`
}

type OutBoundInventoryPayload struct {
	InventoryID    string    `json:"inventory_id"`
	QuantityChange int       `json:"quantity_change"`
	Timestamp      time.Time `json:"timestamp"`
}

func OutBoundInventoryEndpoint(c *fiber.Ctx, db *gorm.DB) error {
	var req OutBoundInventoryRequest
	validate := validator.New()
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": " Validation failed: " + err.Error()})
	}
	inventory_id := c.Params("inventory_id")

	event_payload, _ := json.Marshal(OutBoundInventoryPayload{
		InventoryID:    inventory_id,
		QuantityChange: req.QuantityChange,
		Timestamp:      time.Now(),
	})

	event := model.Event_Stores{
		AggregateID:  &inventory_id,
		EventType:    "InventoryOutBoundQuantityEvent",
		EventPayload: string(event_payload),
	}

	result := db.Create(&event)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create event"})
	}

	topic := event.EventType
	infra.PublishToKafka(topic, event_payload)

	log.Println("✅ InventoryOutBoundQuantityEvent published to Kafka successfully!")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":                   "Inventory event stored and published successfully",
		"inventory_id":              event.EventID,
		"inventory_quantity_change": req.QuantityChange,
	})
}
