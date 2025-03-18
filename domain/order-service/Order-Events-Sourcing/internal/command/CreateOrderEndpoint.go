package command

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/ORDER-EVENTS-SOURCING/internal/infra"
	"github.com/MD-PROJECT/ORDER-EVENTS-SOURCING/internal/model"
)

type CreateOrderRequest struct {
	StoreID       string  `json:"store_id" validate:"required"`
	CustomerID    string  `json:"customer_id" validate:"required"`
	ProductID     string  `json:"product_id" validate:"required"`
	OrderPrice    float64 `json:"order_price" validate:"required,gt=0"`
	OrderQuantity int     `json:"order_quantity" validate:"required,gt=0"`
	OrderDetails  string  `json:"order_details" validate:"required"`
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

func CreateOrderEndpoint(c *fiber.Ctx, db *gorm.DB) error {
	var aggregateID = generateUniqueUUID(db)
	var req CreateOrderRequest
	validate := validator.New()
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": " Validation failed: " + err.Error()})
	}

	event_payload, _ := json.Marshal(model.Orders{
		OrderID:       aggregateID,
		StoreID:       req.StoreID,
		CustomerID:    req.CustomerID,
		ProductID:     req.ProductID,
		OrderStatus:   "waiting",
		OrderPrice:    req.OrderPrice,
		OrderQuantity: req.OrderQuantity,
		OrderDetails:  req.OrderDetails,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	})

	event := model.Event_Stores{
		AggregateID:  &aggregateID,
		EventType:    "OrderCreatedEvent",
		EventPayload: string(event_payload),
	}

	result := db.Create(&event)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create event"})
	}

	topic := event.EventType
	infra.PublishToKafka(topic, event_payload)

	log.Println("✅ OrderCreatedEvent published to Kafka successfully!")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Order event stored and published successfully",
		"order_id": event.EventID,
		"order":    event.EventPayload,
	})
}

// package command

// import (
// 	"encoding/json"
// 	"log"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"gorm.io/gorm"

// 	"github.com/MD-PROJECT/ORDER-EVENTS-SOURCING/internal/infra"
// 	"github.com/MD-PROJECT/ORDER-EVENTS-SOURCING/internal/model"
// 	"github.com/MD-PROJECT/ORDER-EVENTS-SOURCING/internal/utils"
// )

// type CreateOrderRequest struct {
// 	StoreID       string  `json:"store_id"
//
// :"required"`
// 	CustomerID    string  `json:"customer_id" validate:"required"`
// 	OrderPrice    float64 `json:"order_price" validate:"required,gt=0"`
// 	OrderQuantity int     `json:"order_quantity" validate:"required,gt=0"`
// 	OrderDetails  string  `json:"order_details" validate:"required"`
// }

// type OrderCreatedEventPayload struct {
// 	StoreID       string    `json:"store_id"`
// 	CustomerID    string    `json:"customer_id"`
// 	OrderStatus   string    `json:"order_status"`
// 	OrderPrice    float64   `json:"order_price"`
// 	OrderQuantity int       `json:"order_quantity"`
// 	OrderDetails  string    `json:"order_details"`
// 	Timestamp     time.Time `json:"timestamp"`
// 	OrderID       string    `json:"order_id,omitempty"`
// }

// func CreateOrderEndpoint(c *fiber.Ctx, db *gorm.DB) error {
// 	var req CreateOrderRequest
// 	if err := utils.ParseAndValidateRequest(c, &req); err != nil {
// 		return err
// 	}

// 	// เริ่ม transaction
// 	err := db.Transaction(func(tx *gorm.DB) error {
// 		// สร้าง event payload
// 		eventPayloadStruct := OrderCreatedEventPayload{
// 			StoreID:       req.StoreID,
// 			CustomerID:    req.CustomerID,
// 			OrderStatus:   "waiting",
// 			OrderPrice:    req.OrderPrice,
// 			OrderQuantity: req.OrderQuantity,
// 			OrderDetails:  req.OrderDetails,
// 			Timestamp:     time.Now(),
// 		}

// 		payloadBytes, err := json.Marshal(eventPayloadStruct)
// 		if err != nil {
// 			log.Printf("Error marshaling event payload: %v", err)
// 			return err
// 		}

// 		// บันทึก event ลง event store // จะ rollback อัตโนมัติ
// 		eventRecord := model.Event{
// 			EventType:    "OrderCreatedEvent",
// 			EventPayload: string(payloadBytes),
// 		}
// 		if err := tx.Create(&eventRecord).Error; err != nil {
// 			log.Printf("Failed to save event: %v", err)
// 			return err
// 		}

// 		// เพิ่ม order_id เข้า payload ก่อนส่งไป Kafka // จะ rollback อัตโนมัติ
// 		eventPayloadStruct.OrderID = *eventRecord.AggregateID
// 		payloadBytes, err = json.Marshal(eventPayloadStruct)
// 		if err != nil {
// 			log.Printf("Error marshaling event payload with order_id: %v", err)
// 			return err // จะ rollback อัตโนมัติ
// 		}

// 		// ส่งไป Kafka Broker // จะ rollback อัตโนมัติ
// 		topic := eventRecord.EventType
// 		if err := infra.PublishToKafka(topic, payloadBytes); err != nil {
// 			log.Printf("Failed to publish to Kafka: %v", err)
// 			return err
// 		}

// 		log.Println("✅ OrderCreatedEvent published to Kafka successfully!")
// 		return nil // commit transaction
// 	})

// 	// ตรวจสอบผลลัพธ์จาก transaction
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process order creation: " + err.Error()})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": "Order event stored and published successfully",
// 	})
// }
