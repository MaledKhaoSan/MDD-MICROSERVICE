package OrderProducer

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MD-PROJECT/PUBLISHER-SERVICE/internal/infra"
	"github.com/gofiber/fiber/v2"
)

type OrderUpdateStatusEventModel struct {
	OrderID       string    `json:"order_id" validate:"required"`
	CustomerID    string    `json:"customer_id" validate:"required"`
	StoreID       string    `json:"store_id" validate:"required"`
	OrderStatus   string    `json:"order_status" validate:"required"`
	OrderPrice    float64   `json:"order_price" validate:"gte=0,required"`
	OrderQuantity int       `json:"order_quantity" validate:"gt=0,required"`
	OrderDetails  string    `json:"order_details" validate:"omitempty"`
	CreatedAt     time.Time `json:"created_at" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at" validate:"required"`
}

func PublishOrderUpdateHandler(c *fiber.Ctx) error {
	var update OrderUpdateStatusEventModel
	if err := c.BodyParser(&update); err != nil {
		log.Printf("❌ Invalid JSON for order update: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	payload, err := json.Marshal(update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to marshal order update"})
	}

	// ส่งไปยัง topic "OrderUpdatedStatusEvent"
	if err := infra.PublishToKafka("OrderUpdatedStatusEvent", payload); err != nil {
		log.Printf("❌ Failed to publish order update event: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to publish order update event"})
	}

	log.Printf("📤 Sending Kafka Payload: %s", string(payload))
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Order update event published"})
}
