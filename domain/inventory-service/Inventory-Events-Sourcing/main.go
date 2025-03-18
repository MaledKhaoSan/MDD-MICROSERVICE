package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/INVENTORY-EVENTS-SOURCING/internal/infra"
	"github.com/MD-PROJECT/INVENTORY-EVENTS-SOURCING/internal/router"
)

func main() {
	// 1) โหลดค่า .env
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2) อ่าน DATABASE_URL จาก Environment
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// 3) เชื่อมต่อกับ PostgreSQL ผ่าน GORM
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("✅ Successfully connected to PostgreSQL!")

	// 4) เริ่มต้น Kafka Producer
	infra.InitKafkaProducer()
	defer infra.CloseKafkaProducer()

	// 5) สร้าง Fiber App
	app := fiber.New()

	// 6) ตั้งค่า Routes และส่ง `db` เข้าไป
	router.SetupRoutes(app, db)

	// 7) Start Server
	log.Fatal(app.Listen(":8183"))
}
