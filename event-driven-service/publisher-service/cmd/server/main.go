// publisher-service/cmd/server/main.go
package main

import (
	"log"

	OrderProducer "github.com/MD-PROJECT/PUBLISHER-SERVICE/internal/app/order-producer"
	"github.com/MD-PROJECT/PUBLISHER-SERVICE/internal/infra"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Starting Publisher Service...")

	// เริ่มต้น Kafka producer
	infra.InitKafkaProducer()
	defer infra.CloseKafkaProducer() // ปิด producer เมื่อโปรแกรมจบ

	appFiber := fiber.New()

	appFiber.Post("/producer/orderService/publishOrderCreatedEvent", OrderProducer.PublishOrderHandler)
	appFiber.Post("/producer/orderService/publishOrderUpdatedStatusEvent", OrderProducer.PublishOrderUpdateHandler)

	if err := appFiber.Listen(":8081"); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
