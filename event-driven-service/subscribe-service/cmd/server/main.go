package main

import (
	"log"

	"github.com/MD-PROJECT/SUBSCRIBE-SERVICE/internal/infra"
)

func main() {
	log.Println("🚀 Starting Subscribe-Service...")

	// ✅ Start Kafka Consumer (Background Goroutine)
	go infra.StartKafkaConsumer()

	// ✅ Subscribe Service ไม่มี API (เป็น Worker) → รอรับ Event อย่างเดียว
	select {} // บล็อก main goroutine ไว้เพื่อไม่ให้โปรแกรมปิด
}
