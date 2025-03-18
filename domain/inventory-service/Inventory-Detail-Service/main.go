package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/INVENTORY-DETAIL-SERVICE/internal/infra"
	"github.com/MD-PROJECT/INVENTORY-DETAIL-SERVICE/internal/router"
)

func main() {
	// 1️⃣ โหลดค่า .env
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2️⃣ อ่าน DATABASE_URL จาก Environment
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// 3️⃣ เชื่อมต่อกับ PostgreSQL ผ่าน GORM
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("✅ Successfully connected to PostgreSQL!")

	// 4️⃣ สร้าง Fiber App
	app := fiber.New()

	// 5️⃣ ตั้งค่า Routes และส่ง `db` เข้าไป
	router.SetupRoutes(app, db)

	// เริ่ม Kafka Consumer (Background Process)
	go infra.StartKafkaConsumer(db)

	// 6️⃣ Start Server
	log.Fatal(app.Listen(":8185"))
}
