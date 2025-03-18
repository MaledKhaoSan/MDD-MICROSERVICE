package query

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/ORDER-HISTORY-SERVICE/internal/model"
)

// GetOrderDetailsEndpoint ดึงข้อมูล order ตาม ID (Sync)
func GetOrderDetailsEndpoint(c *fiber.Ctx, db *gorm.DB) error {

	orderID := c.Params("id")

	// ค้นหา Order ใน Database
	var order model.Orders
	if err := db.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		// ถ้าไม่พบข้อมูล
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
		}
		// กรณีเกิด Error อื่น
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	log.Printf("✅ Get Order Detail: \n%+v\n", order)
	return c.Status(fiber.StatusOK).JSON(order)
}
