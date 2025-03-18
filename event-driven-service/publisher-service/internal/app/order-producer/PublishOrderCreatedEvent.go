package OrderProducer

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MD-PROJECT/PUBLISHER-SERVICE/internal/infra"
	"github.com/gofiber/fiber/v2"
)

type OrderCreatedEventModel struct {
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

// PublishOrderHandler จัดการคำขอ POST เพื่อ publish order created event
func PublishOrderHandler(c *fiber.Ctx) error {
	var order OrderCreatedEventModel
	if err := c.BodyParser(&order); err != nil {
		log.Printf("❌ Invalid JSON for order: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	payload, err := json.Marshal(order)
	if err != nil {
		log.Printf("❌ Failed to marshal order: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to marshal order"})
	}

	// ส่งไปยัง topic "OrderCreatedEvent"
	if err := infra.PublishToKafka("OrderCreatedEvent", payload); err != nil {
		log.Printf("❌ Failed to publish order created event: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to publish order created event"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Order created event published",
		"order":   order,
	})
}
