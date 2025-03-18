package command

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MD-PROJECT/ORDER-EVENTS-SOURCING/internal/infra"
	"github.com/MD-PROJECT/ORDER-EVENTS-SOURCING/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UpdateOrderStatusRequest struct {
	OrderStatus string `json:"order_status" validate:"required"`
}


type UpdateOrderStatusPayload struct {
	OrderID     string    `json:"order_id"`
	OrderStatus string    `json:"order_status"`
	Timestamp   time.Time `json:"timestamp"`
}

func UpdateOrderStatusEndpoint(c *fiber.Ctx, db *gorm.DB) error {
	var req UpdateOrderStatusRequest
	validate := validator.New()
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": " Validation failed: " + err.Error()})
	}
	orderID := c.Params("orderID")

	event_payload, _ := json.Marshal(UpdateOrderStatusPayload{
		OrderID:     orderID,
		OrderStatus: req.OrderStatus,
		Timestamp:   time.Now(),
	})

	event := model.Event_Stores{
		AggregateID:  &orderID,
		EventType:    "OrderUpdateStatusEvent",
		EventPayload: string(event_payload),
	}

	result := db.Create(&event)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create event"})
	}

	topic := event.EventType
	infra.PublishToKafka(topic, event_payload)

	log.Println("✅ OrderUpdateStatusEvent published to Kafka successfully!")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Order event stored and published successfully",
		"order_id":     event.EventID,
		"order_status": event.EventPayload,
	})
}
