package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MD-PROJECT/INVENTORY-EVENTS-SOURCING/internal/command"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Patch("/inv/:inventory_id/inbound", func(c *fiber.Ctx) error {
		return command.InBoundInventoryEndpoint(c, db)
	})

	app.Patch("/inv/:inventory_id/outbound", func(c *fiber.Ctx) error {
		return command.OutBoundInventoryEndpoint(c, db)
	})

}
