package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/ORDER-EVENTS-SOURCING/internal/command"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// สร้าง Order
	app.Post("/orders/", func(c *fiber.Ctx) error {
		return command.CreateOrderEndpoint(c, db)
	})

	// อัปเดตสถานะ Order
	app.Patch("/orders/:orderID/status", func(c *fiber.Ctx) error {
		return command.UpdateOrderStatusEndpoint(c, db)
	})
}
